package merchant

// import (
// 	"net/http"

// 	db "github.com/akmal4410/gestapo/pkg/database"
// 	"github.com/akmal4410/gestapo/pkg/helpers"
// 	"github.com/akmal4410/gestapo/pkg/helpers/logger"
// 	s3 "github.com/akmal4410/gestapo/pkg/service/s3_service"
// )

// const (
// 	InternalServerError string = "Internal server error"
// 	InvalidBody         string = "Invalid Body"
// 	StatusBadRequest    string = "Status Bad Request"
// )

// type MerchantHandler struct {
// 	log       logger.Logger
// 	storage   *database.MarchantStore
// 	dbStorage *db.DBStore
// 	s3Service *s3.S3Service
// }

// func NewMerchentHandler(logger logger.Logger, storage *database.MarchantStore, dbStorage *db.DBStore, s3Service *s3.S3Service) *MerchantHandler {
// 	return &MerchantHandler{
// 		log:       logger,
// 		storage:   storage,
// 		dbStorage: dbStorage,
// 		s3Service: s3Service,
// 	}
// }

// func (handler *MerchantHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]

// 	product, err := handler.dbStorage.GetProductById(productId)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			handler.log.LogError("Error while GetProductById Not found", err)
// 			helpers.ErrorJson(http.StatusNotFound, "Not found")
// 			return
// 		}
// 		handler.log.LogError("Error while GetProductById", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}

// 	if product.ProductImages != nil {
// 		for i, image := range product.ProductImages {
// 			url, err := handler.s3Service.GetPreSignedURL(image)
// 			if err != nil {
// 				handler.log.LogError("Error while GetPreSignedURL", err)
// 				helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 				return
// 			}
// 			product.ProductImages[i] = url
// 		}
// 	}
// 	helpers.WriteJSON(w, http.StatusOK, product)
// }

// func (handler *MerchantHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	product, err := handler.dbStorage.GetProductById(productId)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			handler.log.LogError("Error while GetProductById Not fount", err)
// 			helpers.ErrorJson(http.StatusNotFound, "Not found")
// 			return
// 		}
// 		handler.log.LogError("Error while retrieving product", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}

// 	payload, ok := r.Context().Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)
// 	if !ok {
// 		err := errors.New("unable to retrieve user payload from context")
// 		handler.log.LogError("Error", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}

// 	if product.MerchantID != payload.UserID {
// 		err := errors.New("unauthorized: product does not belong to the authenticated merchant")
// 		handler.log.LogError("Error", err)
// 		helpers.ErrorJson(http.StatusForbidden, err.Error())
// 		return
// 	}

// 	err = handler.storage.DeleteProduct(productId)
// 	if err != nil {
// 		handler.log.LogError("Error while DeleteProduct", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}

// 	for _, key := range product.ProductImages {
// 		err := handler.s3Service.DeleteKey(key)
// 		if err != nil {
// 			handler.log.LogError("Error deleting file from S3", err)
// 			helpers.ErrorJson(http.StatusInternalServerError, "Error deleting file from")
// 			return
// 		}
// 	}
// 	helpers.WriteJSON(w, http.StatusOK, "Product deleted succesfully")
// }

// func (handler *MerchantHandler) AddProductDiscount(w http.ResponseWriter, r *http.Request) {
// 	req := new(entity.AddDiscountReq)
// 	err := helpers.ValidateBody(r.Body, req)
// 	if err != nil {
// 		handler.log.LogError("Error while ValidateBody", err)
// 		helpers.ErrorJson(http.StatusBadRequest, InvalidBody)
// 		return
// 	}
// 	_, err = handler.dbStorage.GetProductById(req.ProductId)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			handler.log.LogError("Error while GetProductById Not fount", err)
// 			helpers.ErrorJson(http.StatusNotFound, "Product Not found")
// 			return
// 		}
// 		handler.log.LogError("Error while retrieving product", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}

// 	err = handler.storage.AddProductDiscount(req)
// 	if err != nil {
// 		handler.log.LogError("Error while ApplyProductDiscount", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}

// 	helpers.WriteJSON(w, http.StatusOK, "Discount added successfully")
// }

// func (handler *MerchantHandler) EditProductDiscount(w http.ResponseWriter, r *http.Request) {

// 	req := new(entity.EditDiscountReq)
// 	err := helpers.ValidateBody(r.Body, req)
// 	if err != nil {
// 		handler.log.LogError("Error while ValidateBody", err)
// 		helpers.ErrorJson(http.StatusBadRequest, InvalidBody)
// 		return
// 	}

// 	id := mux.Vars(r)["id"]
// 	res, err := handler.storage.CheckDataExist("discounts", "id", id)
// 	if err != nil {
// 		handler.log.LogError("Error while CheckDataExist", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}
// 	if !res {
// 		err = fmt.Errorf("discounts doesnt exist: %s", id)
// 		handler.log.LogError("Error ", err)
// 		helpers.ErrorJson(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	err = handler.storage.EditProductDiscount(id, req)
// 	if err != nil {
// 		handler.log.LogError("Error while ApplyProductDiscount", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}
// 	helpers.WriteJSON(w, http.StatusOK, "Discount updated successfully")
// }
