package controllers

import (
	"fmt"
	"net/http"

	"github.com/akmal4410/gestapo/database"
	"github.com/akmal4410/gestapo/helpers"
	"github.com/akmal4410/gestapo/models"
	"github.com/akmal4410/gestapo/services/mail"
	"github.com/akmal4410/gestapo/services/twilio"
	"github.com/akmal4410/gestapo/utils"
)

type AuthController struct {
	twilioService twilio.TwilioService
	emailService  mail.EmailService
	storage       *database.Storage
}

func NewAuthController(twilio twilio.TwilioService, email mail.EmailService, storage *database.Storage) *AuthController {
	return &AuthController{
		twilioService: twilio,
		emailService:  email,
		storage:       storage,
	}
}

func (auth AuthController) SendOTP(w http.ResponseWriter, r *http.Request) {
	req := new(models.SendOTPReq)

	err := helpers.ValidateBody(r, req)
	if err != nil {
		helpers.ErrorJson(w, http.StatusBadRequest, err)
		return
	}

	err = utils.ValidateEmailOrPhone(req)
	if err != nil {
		helpers.ErrorJson(w, http.StatusBadRequest, err)
		return
	}

	column, value := utils.IdentifiesColumnValue(req.Email, req.Phone)
	if len(column) == 0 {
		helpers.ErrorJson(w, http.StatusBadRequest, err)
		return
	}
	res, err := auth.storage.CheckUserExist(column, value)
	if err != nil {
		helpers.ErrorJson(w, http.StatusInternalServerError, err)
		return
	}
	if res {
		helpers.ErrorJson(w, http.StatusInternalServerError, fmt.Errorf("account already exist using this %s", column))
		return
	}

	if !utils.IsEmpty(req.Email) {
		// res, err := auth.storage.CheckUserExist("email", req.Email)
		// if err != nil {
		// 	helpers.ErrorJson(w, http.StatusInternalServerError, err)
		// 	return
		// }
		// if res {
		// 	helpers.ErrorJson(w, http.StatusInternalServerError, fmt.Errorf("account already exist using this Email"))
		// 	return
		// }
		err = auth.emailService.SendEmail(req.Email, "Sign Up OTP", "Welcome to Gestapo !!!. Use the following OTP to complete your Sign Up procedures. OTP is valid for 5minutes", nil, nil, nil)
		if err != nil {
			helpers.ErrorJson(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		// res, err := auth.storage.CheckUserExist("phone", req.Phone)
		// if err != nil {
		// 	helpers.ErrorJson(w, http.StatusInternalServerError, err)
		// 	return
		// }
		// if res {
		// 	helpers.ErrorJson(w, http.StatusInternalServerError, fmt.Errorf("account already exist using this Phone number"))
		// 	return
		// }
		err = auth.twilioService.SendOTP(req.Phone)
		if err != nil {
			helpers.ErrorJson(w, http.StatusInternalServerError, err)
			return
		}
	}
	w.Header().Set("session-token", "akmal token")

	helpers.WriteJSON(w, http.StatusOK, "OTP sent successfully")
}

func (auth AuthController) LoginUser(w http.ResponseWriter, r *http.Request) {
}

func (auth AuthController) CreateUser(w http.ResponseWriter, r *http.Request) {
}
