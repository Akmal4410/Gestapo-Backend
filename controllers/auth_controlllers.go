package controllers

import (
	"net/http"
)

type Auth struct {
}

func (auth Auth) LoginUser(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusAccepted, "You are tring to login")
}

func (auth Auth) CreateUser(w http.ResponseWriter, r *http.Request) {
}
