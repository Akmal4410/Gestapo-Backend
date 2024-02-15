package controllers

import (
	"net/http"

	"github.com/akmal4410/gestapo/helpers"
	"github.com/akmal4410/gestapo/models"
	"github.com/akmal4410/gestapo/services/twilio"
)

type Auth struct {
	twilio twilio.OTPService
}

func (auth Auth) SendOTP(w http.ResponseWriter, r *http.Request) {
	payload := new(models.SendOTPReq)

	err := helpers.ValidateBody(r, payload)
	if err != nil {
		helpers.ErrorJson(w, http.StatusBadRequest, err)
		return
	}

	err = auth.twilio.SendOTP(payload.Phone)
	if err != nil {
		helpers.ErrorJson(w, http.StatusInternalServerError, err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, "")
}

func (auth Auth) LoginUser(w http.ResponseWriter, r *http.Request) {
}

func (auth Auth) CreateUser(w http.ResponseWriter, r *http.Request) {
}
