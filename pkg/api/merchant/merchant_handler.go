package merchant

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/akmal4410/gestapo/pkg/api/merchant/database"
	"github.com/akmal4410/gestapo/pkg/api/merchant/database/entity"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/service/logger"
	s3 "github.com/akmal4410/gestapo/pkg/service/s3_service"
	"github.com/akmal4410/gestapo/pkg/service/token"
	"github.com/akmal4410/gestapo/pkg/utils"
	"github.com/gorilla/mux"
)

const (
	InternalServerError string = "Internal server error"
	InvalidBody         string = "Invalid Body"
	StatusBadRequest    string = "Status Bad Request"
)

type MerchantHandler struct {
	storage   *database.MarchantStore
	log       logger.Logger
	s3Service *s3.S3Service
}

func NewMerchentHandler(storage *database.MarchantStore, logger logger.Logger, s3Service *s3.S3Service) *MerchantHandler {
	return &MerchantHandler{
		storage:   storage,
		log:       logger,
		s3Service: s3Service,
	}
}

func (handler *MerchantHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	res, err := handler.storage.CheckUserExist("id", userId)
	if err != nil {
		handler.log.LogError("Error while CheckUserExist", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	if !res {
		err = fmt.Errorf("account does'nt exist using %s", userId)
		handler.log.LogError(err)
		helpers.ErrorJson(w, http.StatusConflict, err.Error())
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

	res, err := handler.storage.CheckUserExist("id", payload.UserID)
	if err != nil {
		handler.log.LogError("Error while CheckUserExist", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	if !res {
		err = fmt.Errorf("account does'nt exist using %s", payload.UserID)
		handler.log.LogError(err)
		helpers.ErrorJson(w, http.StatusConflict, err.Error())
		return
	}

	files := r.MultipartForm.File["files"]
	if len(files) > maxFileCount {
		handler.log.LogError("Too many files uploaded", "Max allowed: %d", maxFileCount)
		errMsg := fmt.Sprintf("too many files uploaded. Max allowed: %s", strconv.Itoa(maxFileCount))
		helpers.ErrorJson(w, http.StatusBadRequest, errMsg)
		return
	}

	var uploadedFileURLs []string
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			handler.log.LogError("Unable to open file", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, "Unable to open file")
			return
		}
		defer file.Close()

		// Upload file to S3

		folderPath := "profile/" + payload.UserID + "/"
		fileURL, err := handler.s3Service.UploadFileToS3(file, folderPath, fileHeader.Filename)
		if err != nil {
			handler.log.LogError("Error uploading file to S3", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, "Error uploading file to S3")
			return
		}

		handler.log.LogInfo("File uploaded to S3 successfully", "FileURL:", fileURL)
		uploadedFileURLs = append(uploadedFileURLs, fileURL)
	}

	if len(uploadedFileURLs) != 0 {
		req.ProfileImage = uploadedFileURLs[0]
	}
	err = handler.storage.UpdateProfile(payload.UserID, req)
	if err != nil {
		handler.log.LogError("Error while UpdateProfile", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "User updated successfully")
}
