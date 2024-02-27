package merchant

import (
	"fmt"
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/merchant/database"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/service/logger"
	"github.com/gorilla/mux"
)

const (
	InternalServerError string = "Internal server error"
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

func (merchat *MerchantHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	user_id := mux.Vars(r)["id"]

	res, err := merchat.storage.CheckUserExist("id", user_id)
	if err != nil {
		merchat.log.LogError("Error while CheckUserExist", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	if !res {
		err = fmt.Errorf("account does'nt exist using this %s", user_id)
		merchat.log.LogError(err)
		helpers.ErrorJson(w, http.StatusConflict, err.Error())
		return
	}

}
