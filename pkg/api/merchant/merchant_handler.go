package merchant

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/akmal4410/gestapo/pkg/api/merchant/database"
	"github.com/akmal4410/gestapo/pkg/api/merchant/database/entity"
	db "github.com/akmal4410/gestapo/pkg/database"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
	"github.com/akmal4410/gestapo/pkg/helpers/token"
	s3 "github.com/akmal4410/gestapo/pkg/service/s3_service"
	"github.com/akmal4410/gestapo/pkg/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const (
	InternalServerError string = "Internal server error"
	InvalidBody         string = "Invalid Body"
	StatusBadRequest    string = "Status Bad Request"
)

type MerchantHandler struct {
	log       logger.Logger
	storage   *database.MarchantStore
	dbStorage *db.DBStore
	s3Service *s3.S3Service
}

func NewMerchentHandler(logger logger.Logger, storage *database.MarchantStore, dbStorage *db.DBStore, s3Service *s3.S3Service) *MerchantHandler {
	return &MerchantHandler{
		log:       logger,
		storage:   storage,
		dbStorage: dbStorage,
		s3Service: s3Service,
	}
}

func (handler *MerchantHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	res, err := handler.storage.CheckDataExist("user_data", "id", userId)
	if err != nil {
		handler.log.LogError("Error while CheckUserExist", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	if !res {
		err = fmt.Errorf("account does'nt exist using %s", userId)
		handler.log.LogError(err)
		helpers.ErrorJson(w, http.StatusNotFound, err.Error())
		return
	}

	userData, err := handler.storage.GetProfile(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			handler.log.LogError("Error while GetProfile", err)
			helpers.ErrorJson(w, http.StatusNotFound, "Not found")
			return
		}
		handler.log.LogError("Error while GetProfile", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	if userData.ProfileImage != nil {
		url, err := handler.s3Service.GetPreSignedURL(*userData.ProfileImage)
		if err != nil {
			handler.log.LogError("Error while GetPreSignedURL", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return
		}
		userData.ProfileImage = &url
	}

	helpers.WriteJSON(w, http.StatusOK, userData)
}

func (handler *MerchantHandler) EditProfile(w http.ResponseWriter, r *http.Request) {
	const (
		thirtyTwoMB      = 32 << 20
		maxFileCount int = 1
	)

	// Extract the JSON data from the form
	jsonData := r.FormValue("data")
	reader := io.Reader(strings.NewReader(jsonData))

	req := new(entity.EditMerchantReq)
	err := helpers.ValidateBody(reader, req)
	if err != nil {
		handler.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, InvalidBody)
		return
	}

	err = r.ParseMultipartForm(thirtyTwoMB)
	if err != nil {
		handler.log.LogError("Unable to parse form", err.Error())
		helpers.ErrorJson(w, http.StatusBadRequest, StatusBadRequest)
		return
	}

	payload := r.Context().Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)

	files := r.MultipartForm.File["files"]
	if len(files) > maxFileCount {
		handler.log.LogError("Too many files uploaded", "Max allowed: %d", maxFileCount)
		errMsg := fmt.Sprintf("too many files uploaded. Max allowed: %s", strconv.Itoa(maxFileCount))
		helpers.ErrorJson(w, http.StatusBadRequest, errMsg)
		return
	}

	var uploadedFileKeys []string
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			handler.log.LogError("Unable to open file", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, "Unable to open file")
			return
		}
		defer file.Close()

		folderPath := "profile/" + payload.UserID + "/"
		fileURL, err := handler.s3Service.UploadFileToS3(file, folderPath, fileHeader.Filename)
		if err != nil {
			handler.log.LogError("Error uploading file to S3", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, "Error uploading file to S3")
			return
		}

		handler.log.LogInfo("File uploaded to S3 successfully", "FileURL:", fileURL)
		uploadedFileKeys = append(uploadedFileKeys, fileURL)
	}

	if len(uploadedFileKeys) != 0 {
		req.ProfileImage = uploadedFileKeys[0]
	}
	err = handler.storage.UpdateProfile(payload.UserID, req)
	if err != nil {
		handler.log.LogError("Error while UpdateProfile", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "User updated successfully")
}

func (handler *MerchantHandler) InsertProduct(w http.ResponseWriter, r *http.Request) {
	const (
		thirtyTwoMB      = 32 << 20
		maxFileCount int = 5
	)
	// Extract the JSON data from the form
	jsonData := r.FormValue("data")
	reader := io.Reader(strings.NewReader(jsonData))

	req := new(entity.AddProductReq)
	err := helpers.ValidateBody(reader, req)
	if err != nil {
		handler.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := handler.storage.CheckDataExist("categories", "id", req.CategoryId)
	if err != nil {
		handler.log.LogError("Error while CheckCategoryExist", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	if !res {
		err = fmt.Errorf("category doesnt exist: %s", req.CategoryId)
		handler.log.LogError("Error ", err)
		helpers.ErrorJson(w, http.StatusNotFound, err.Error())
		return
	}

	err = r.ParseMultipartForm(thirtyTwoMB)
	if err != nil {
		handler.log.LogError("Unable to parse form", err.Error())
		helpers.ErrorJson(w, http.StatusBadRequest, StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		handler.log.LogError("There should be atleast one image")
		errMsg := "There should be atleast one image"
		helpers.ErrorJson(w, http.StatusBadRequest, errMsg)
		return
	}
	if len(files) > maxFileCount {
		handler.log.LogError("Too many files uploaded", "Max allowed: %d", maxFileCount)
		errMsg := fmt.Sprintf("too many files uploaded. Max allowed: %s", strconv.Itoa(maxFileCount))
		helpers.ErrorJson(w, http.StatusBadRequest, errMsg)
		return
	}

	payload, ok := r.Context().Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)
	if !ok {
		err := errors.New("unable to retrieve user payload from context")
		handler.log.LogError("Error", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	uuId, err := uuid.NewRandom()
	if err != nil {
		handler.log.LogError("error while uuid NewRandom", err.Error())
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	var uploadedFileKeys []string
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			handler.log.LogError("Unable to open file", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, "Unable to open file")
			return
		}
		defer file.Close()

		folderPath := filepath.Join("products", payload.UserID, uuId.String()) + "/"

		fileURL, err := handler.s3Service.UploadFileToS3(file, folderPath, fileHeader.Filename)
		if err != nil {
			handler.log.LogError("Error uploading file to S3", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, "Error uploading file to S3")
			return
		}

		handler.log.LogInfo("File uploaded to S3 successfully", "FileURL:", fileURL)
		uploadedFileKeys = append(uploadedFileKeys, fileURL)
	}
	if len(uploadedFileKeys) != 0 {
		req.ProductImages = uploadedFileKeys
	}

	err = handler.storage.InsertProduct(payload.UserID, uuId.String(), req)
	if err != nil {
		handler.log.LogError("Error while InsertProduct", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "Product added successfully")
}

func (handler *MerchantHandler) EditProduct(w http.ResponseWriter, r *http.Request) {
	const (
		thirtyTwoMB      = 32 << 20
		maxFileCount int = 5
	)
	// Extract the JSON data from the form
	jsonData := r.FormValue("data")
	reader := io.Reader(strings.NewReader(jsonData))

	req := new(entity.EditProductReq)
	err := helpers.ValidateBody(reader, req)
	if err != nil {
		handler.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, err.Error())
		return
	}

	payload, ok := r.Context().Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)
	if !ok {
		err := errors.New("unable to retrieve user payload from context")
		handler.log.LogError("Error", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	id := mux.Vars(r)["id"]
	product, err := handler.dbStorage.GetProductById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			handler.log.LogError("Error while GetProductById Not fount", err)
			helpers.ErrorJson(w, http.StatusNotFound, "Product Not found")
			return
		}
		handler.log.LogError("Error while retrieving product", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	if product.MerchantID != payload.UserID {
		err := errors.New("unauthorized: product does not belong to the authenticated merchant")
		handler.log.LogError("Error", err)
		helpers.ErrorJson(w, http.StatusForbidden, err.Error())
		return
	}

	if req.ClearImages {
		for _, key := range product.ProductImages {
			err := handler.s3Service.DeleteKey(key)
			if err != nil {
				handler.log.LogError("Error deleting file from S3", err)
				helpers.ErrorJson(w, http.StatusInternalServerError, "Error deleting file from")
				return
			}
		}
		product.ProductImages = make([]string, 0)
	}

	err = r.ParseMultipartForm(thirtyTwoMB)
	if err != nil {
		handler.log.LogError("Unable to parse form", err.Error())
		helpers.ErrorJson(w, http.StatusBadRequest, StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		handler.log.LogError("There should be atleast one image")
		errMsg := "There should be atleast one image"
		helpers.ErrorJson(w, http.StatusBadRequest, errMsg)
		return
	}
	if len(files) > maxFileCount {
		handler.log.LogError("Too many files uploaded", "Max allowed: %d", maxFileCount)
		errMsg := fmt.Sprintf("too many files uploaded. Max allowed: %s", strconv.Itoa(maxFileCount))
		helpers.ErrorJson(w, http.StatusBadRequest, errMsg)
		return
	}

	var uploadedFileKeys []string
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			handler.log.LogError("Unable to open file", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, "Unable to open file")
			return
		}
		defer file.Close()

		folderPath := filepath.Join("products", payload.UserID, id) + "/"

		fileURL, err := handler.s3Service.UploadFileToS3(file, folderPath, fileHeader.Filename)
		if err != nil {
			handler.log.LogError("Error uploading file to S3", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, "Error uploading file to S3")
			return
		}

		handler.log.LogInfo("File uploaded to S3 successfully", "FileURL:", fileURL)
		uploadedFileKeys = append(uploadedFileKeys, fileURL)
	}
	uploadedFileKeys = append(uploadedFileKeys, product.ProductImages...)
	if len(uploadedFileKeys) != 0 {
		req.ProductImages = uploadedFileKeys
	}

	err = handler.storage.UpdateProduct(id, req)
	if err != nil {
		handler.log.LogError("Error while UpdateProduct", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "Product updated successfully")
}

func (handler *MerchantHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	res, err := handler.dbStorage.GetProducts()
	if err != nil {
		handler.log.LogError("Error while GetProducts", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	for _, product := range res {
		if product.ProductImages != nil {
			for i, image := range product.ProductImages {
				url, err := handler.s3Service.GetPreSignedURL(image)
				if err != nil {
					handler.log.LogError("Error while GetPreSignedURL", err)
					helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
					return
				}
				product.ProductImages[i] = url
			}
		}
	}

	helpers.WriteJSON(w, http.StatusOK, res)
}

func (handler *MerchantHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]

	product, err := handler.dbStorage.GetProductById(productId)
	if err != nil {
		if err == sql.ErrNoRows {
			handler.log.LogError("Error while GetProductById Not found", err)
			helpers.ErrorJson(w, http.StatusNotFound, "Not found")
			return
		}
		handler.log.LogError("Error while GetProductById", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	if product.ProductImages != nil {
		for i, image := range product.ProductImages {
			url, err := handler.s3Service.GetPreSignedURL(image)
			if err != nil {
				handler.log.LogError("Error while GetPreSignedURL", err)
				helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
				return
			}
			product.ProductImages[i] = url
		}
	}
	helpers.WriteJSON(w, http.StatusOK, product)
}

func (handler *MerchantHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	product, err := handler.dbStorage.GetProductById(productId)
	if err != nil {
		if err == sql.ErrNoRows {
			handler.log.LogError("Error while GetProductById Not fount", err)
			helpers.ErrorJson(w, http.StatusNotFound, "Not found")
			return
		}
		handler.log.LogError("Error while retrieving product", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	payload, ok := r.Context().Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)
	if !ok {
		err := errors.New("unable to retrieve user payload from context")
		handler.log.LogError("Error", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	if product.MerchantID != payload.UserID {
		err := errors.New("unauthorized: product does not belong to the authenticated merchant")
		handler.log.LogError("Error", err)
		helpers.ErrorJson(w, http.StatusForbidden, err.Error())
		return
	}

	err = handler.storage.DeleteProduct(productId)
	if err != nil {
		handler.log.LogError("Error while DeleteProduct", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	for _, key := range product.ProductImages {
		err := handler.s3Service.DeleteKey(key)
		if err != nil {
			handler.log.LogError("Error deleting file from S3", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, "Error deleting file from")
			return
		}
	}
	helpers.WriteJSON(w, http.StatusOK, "Product deleted succesfully")
}

func (handler *MerchantHandler) AddProductDiscount(w http.ResponseWriter, r *http.Request) {
	req := new(entity.AddDiscountReq)
	err := helpers.ValidateBody(r.Body, req)
	if err != nil {
		handler.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, InvalidBody)
		return
	}
	_, err = handler.dbStorage.GetProductById(req.ProductId)
	if err != nil {
		if err == sql.ErrNoRows {
			handler.log.LogError("Error while GetProductById Not fount", err)
			helpers.ErrorJson(w, http.StatusNotFound, "Product Not found")
			return
		}
		handler.log.LogError("Error while retrieving product", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	err = handler.storage.AddProductDiscount(req)
	if err != nil {
		handler.log.LogError("Error while ApplyProductDiscount", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "Discount added successfully")
}

func (handler *MerchantHandler) EditProductDiscount(w http.ResponseWriter, r *http.Request) {

	req := new(entity.EditDiscountReq)
	err := helpers.ValidateBody(r.Body, req)
	if err != nil {
		handler.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, InvalidBody)
		return
	}

	id := mux.Vars(r)["id"]
	res, err := handler.storage.CheckDataExist("discounts", "id", id)
	if err != nil {
		handler.log.LogError("Error while CheckDataExist", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	if !res {
		err = fmt.Errorf("discounts doesnt exist: %s", id)
		handler.log.LogError("Error ", err)
		helpers.ErrorJson(w, http.StatusNotFound, err.Error())
		return
	}

	err = handler.storage.EditProductDiscount(id, req)
	if err != nil {
		handler.log.LogError("Error while ApplyProductDiscount", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, "Discount updated successfully")
}
