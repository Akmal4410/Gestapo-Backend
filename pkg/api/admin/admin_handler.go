package admin

import (
	"fmt"
	"net/http"

	"github.com/akmal4410/gestapo/pkg/grpc_api/admin_service/db"
	"github.com/akmal4410/gestapo/pkg/grpc_api/admin_service/db/entity"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/helpers/logger"
)

const (
	InternalServerError string = "Internal server error"
	InvalidBody         string = "Invalid Body"
)

type AdminHandler struct {
	storage *db.AdminStore
	log     logger.Logger
}

func NewAdminHandler(storage *db.AdminStore, log logger.Logger) *AdminHandler {
	return &AdminHandler{
		storage: storage,
		log:     log,
	}
}

func (handler *AdminHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	req := new(entity.AddCategoryReq)

	err := helpers.ValidateBody(r.Body, req)
	if err != nil {
		handler.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(http.StatusBadRequest, InvalidBody)
		return
	}

	res, err := handler.storage.CheckCategoryExist(req.Category_Name)
	if err != nil {
		handler.log.LogError("Error while CheckCategoryExist", err)
		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
		return
	}
	if res {
		err = fmt.Errorf("category already exist: %s", req.Category_Name)
		handler.log.LogError("Error ", err)
		helpers.ErrorJson(http.StatusNotFound, err.Error())
		return
	}

	err = handler.storage.AddCategory(req)
	if err != nil {
		handler.log.LogError("Error while InsertCategory", err)
		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "Category insterted Successfully")
}

func (handler *AdminHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	res, err := handler.storage.GetCategories()
	if err != nil {
		handler.log.LogError("Error while GetCategories", err)
		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, res)
}

func (handler *AdminHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := handler.storage.GetUsers()
	if err != nil {
		handler.log.LogError("Error while GetUsers", err)
		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, res)
}
