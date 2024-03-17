package user

import (
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/user/database"
	"github.com/akmal4410/gestapo/pkg/api/user/database/entity"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/service/logger"
)

const (
	InternalServerError string = "Internal server error"
	InvalidBody         string = "Invalid Body"
	StatusBadRequest    string = "Status Bad Request"
)

type UserHandler struct {
	log     logger.Logger
	storage *database.UserStore
}

func NewUserHandler(storage *database.UserStore) *UserHandler {
	return &UserHandler{
		storage: storage,
	}
}

func (handler *UserHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	res := new(entity.GetHomeRes)
	var err error
	res.Discount, err = handler.storage.GetDiscount()
	if err != nil {
		handler.log.LogError("Error while GetDiscount", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, res)
}
