package admin

import (
	"fmt"
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/admin/database"
	"github.com/akmal4410/gestapo/pkg/api/admin/database/entity"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/service/logger"
)

const (
	InternalServerError string = "Internal server error"
	InvalidBody         string = "Invalid Body"
)

type AdminHandler struct {
	storage *database.AdminStore
	log     logger.Logger
}

func NewAdminHandler(storage *database.AdminStore, log logger.Logger) *AdminHandler {
	return &AdminHandler{
		storage: storage,
		log:     log,
	}
}

func (handler *AdminHandler) CreateCategories(w http.ResponseWriter, r *http.Request) {
	req := new(entity.InsertCategoryReq)

	err := helpers.ValidateBody(r, req)
	if err != nil {
		handler.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, InvalidBody)
		return
	}

	res, err := handler.storage.CheckCategoryExist(req.Category_Name)
	if err != nil {
		handler.log.LogError("Error while CheckCategoryExist", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	if res {
		err = fmt.Errorf("category already exist: %s", req.Category_Name)
		handler.log.LogError("Error ", err)
		helpers.ErrorJson(w, http.StatusConflict, err.Error())
		return
	}

	err = handler.storage.InsertCategory(req)
	if err != nil {
		handler.log.LogError("Error while InsertCategory", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "Category insterted Successfully")
}
