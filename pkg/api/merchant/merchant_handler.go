package merchant

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/merchant/database"
	"github.com/akmal4410/gestapo/pkg/api/merchant/database/entity"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/service/logger"
	"github.com/akmal4410/gestapo/pkg/service/token"
	"github.com/akmal4410/gestapo/pkg/utils"
	"github.com/gorilla/mux"
)

const (
	InternalServerError string = "Internal server error"
	InvalidBody         string = "Invalid Body"
)

type MerchantHandler struct {
	storage *database.MarchantStore
	log     logger.Logger
}

func NewMerchentHandler(storage *database.MarchantStore, logger logger.Logger) *MerchantHandler {
	return &MerchantHandler{
		storage: storage,
		log:     logger,
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
	helpers.WriteJSON(w, http.StatusOK, userData)
}

func (handler *MerchantHandler) EditProfile(w http.ResponseWriter, r *http.Request) {
	req := new(entity.EditMerchantReq)

	err := helpers.ValidateBody(r, req)
	if err != nil {
		handler.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, InvalidBody)
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

	err = handler.storage.UpdateProfile(payload.UserID, req)
	if err != nil {
		handler.log.LogError("Error while UpdateProfile", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "User updated successfully")
}
