package service

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"

// 	"github.com/akmal4410/gestapo/pkg/helpers"
// 	"github.com/gorilla/mux"
// )

// func (handler *MerchantHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
// 	userId := mux.Vars(r)["id"]

// 	res, err := handler.storage.CheckDataExist("user_data", "id", userId)
// 	if err != nil {
// 		handler.log.LogError("Error while CheckUserExist", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}

// 	if !res {
// 		err = fmt.Errorf("account does'nt exist using %s", userId)
// 		handler.log.LogError(err)
// 		helpers.ErrorJson(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	userData, err := handler.storage.GetProfile(userId)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			handler.log.LogError("Error while GetProfile", err)
// 			helpers.ErrorJson(http.StatusNotFound, "Not found")
// 			return
// 		}
// 		handler.log.LogError("Error while GetProfile", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}

// 	if userData.ProfileImage != nil {
// 		url, err := handler.s3Service.GetPreSignedURL(*userData.ProfileImage)
// 		if err != nil {
// 			handler.log.LogError("Error while GetPreSignedURL", err)
// 			helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 			return
// 		}
// 		userData.ProfileImage = &url
// 	}

// 	helpers.WriteJSON(w, http.StatusOK, userData)
// }
