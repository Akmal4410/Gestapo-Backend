package controller

import (
	"net/http"

	"github.com/akmal4410/gestapo/util"
)

type AuthController struct {
}

func (auth AuthController) LoginUser(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, http.StatusAccepted, "You are tring to login")
}

func (auth AuthController) CreateUser(w http.ResponseWriter, r *http.Request) {
}
