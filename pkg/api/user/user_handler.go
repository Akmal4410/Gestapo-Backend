package user

import (
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/user/database"
	db "github.com/akmal4410/gestapo/pkg/database"
	s3 "github.com/akmal4410/gestapo/pkg/service/s3_service"

	"github.com/akmal4410/gestapo/pkg/api/user/database/entity"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
)

const (
	InternalServerError string = "Internal server error"
	InvalidBody         string = "Invalid Body"
	StatusBadRequest    string = "Status Bad Request"
)

type UserHandler struct {
	log       logger.Logger
	storage   *database.UserStore
	dbStorage *db.DBStore
	s3Service *s3.S3Service
}

func NewUserHandler(log logger.Logger, storage *database.UserStore, dbStorage *db.DBStore, s3Service *s3.S3Service) *UserHandler {
	return &UserHandler{
		log:       log,
		storage:   storage,
		dbStorage: dbStorage,
		s3Service: s3Service,
	}
}

func (handler *UserHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	res := new(entity.GetHomeRes)
	var err error
	res.Discount, err = handler.storage.GetDiscount()
	if err != nil {
		handler.log.LogError("Error while GetDiscount", err)
		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
		return
	}
	//Converting the key to presigned url
	url, err := handler.s3Service.GetPreSignedURL(res.Discount.ProductImage)
	if err != nil {
		handler.log.LogError("Error while GetPreSignedURL", err)
		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
		return
	}
	res.Discount.ProductImage = url

	res.Merchants, err = handler.storage.GetMerchants()
	if err != nil {
		handler.log.LogError("Error while GetMerchants", err)
		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
		return
	}
	for _, merchant := range res.Merchants {
		if merchant.ImageURL != nil {
			url, err := handler.s3Service.GetPreSignedURL(*merchant.ImageURL)
			if err != nil {
				handler.log.LogError("Error while GetPreSignedURL", err)
				helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
				return
			}
			merchant.ImageURL = &url
		}
	}

	res.Products, err = handler.dbStorage.GetProducts()
	if err != nil {
		handler.log.LogError("Error while GetProducts", err)
		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
		return
	}
	for _, product := range res.Products {
		for i, image := range product.ProductImages {
			url, err := handler.s3Service.GetPreSignedURL(image)
			if err != nil {
				handler.log.LogError("Error while GetPreSignedURL", err)
				helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
				return
			}
			product.ProductImages[i] = url
		}

	}
	helpers.WriteJSON(w, http.StatusOK, res)
}
