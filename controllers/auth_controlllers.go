package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/akmal4410/gestapo/database"
	"github.com/akmal4410/gestapo/helpers"
	"github.com/akmal4410/gestapo/middlewares"
	"github.com/akmal4410/gestapo/models"
	"github.com/akmal4410/gestapo/services/cache"
	"github.com/akmal4410/gestapo/services/logger"
	"github.com/akmal4410/gestapo/services/mail"
	"github.com/akmal4410/gestapo/services/token"
	"github.com/akmal4410/gestapo/services/twilio"
	"github.com/akmal4410/gestapo/utils"
	"github.com/pkg/errors"
)

const (
	InternalServerError string = "Internal server error"
	InvalidBody         string = "Invalid Body"
)

type AuthController struct {
	twilioService twilio.TwilioService
	emailService  mail.EmailService
	storage       *database.Storage
	token         token.Maker
	redis         cache.Cache
	log           logger.Logger
}

func NewAuthController(
	twilio twilio.TwilioService,
	email mail.EmailService,
	storage *database.Storage,
	token token.Maker,
	redisCache cache.Cache,
	logger logger.Logger,
) *AuthController {
	return &AuthController{
		twilioService: twilio,
		emailService:  email,
		storage:       storage,
		token:         token,
		redis:         redisCache,
		log:           logger,
	}
}

func (auth *AuthController) SendOTP(w http.ResponseWriter, r *http.Request) {
	req := new(models.SendOTPReq)

	err := helpers.ValidateBody(r, req)
	if err != nil {
		auth.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, InvalidBody)
		return
	}

	err = utils.ValidateEmailOrPhone(req.Email, req.Phone)
	if err != nil {
		auth.log.LogError("Error while ValidateEmailOrPhone", err)
		helpers.ErrorJson(w, http.StatusBadRequest, "Invalid Email or Phone")
		return
	}

	column, value := utils.IdentifiesColumnValue(req.Email, req.Phone)
	if req.Action == utils.SIGN_UP {
		if len(column) == 0 {
			auth.log.LogError("Error while IdentifiesColumnValue", column)
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return
		}
		res, err := auth.storage.CheckDataExist(column, value)
		if err != nil {
			auth.log.LogError("Error while CheckDataExist", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return
		}
		if res {
			err = fmt.Errorf("account already exist using this %s", column)
			auth.log.LogError(err)
			helpers.ErrorJson(w, http.StatusConflict, err.Error())
			return
		}
	}

	if !utils.IsEmpty(req.Email) {
		err = auth.emailService.SendOTP(req.Email, utils.EmailSubject, utils.EmailSubject, auth.redis)
		if err != nil {
			auth.log.LogError("Error while SendOTP", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return
		}
	} else {
		phoneNumber := fmt.Sprintf("+91%s", req.Phone)
		err = auth.twilioService.SendOTP(phoneNumber)
		if err != nil {
			auth.log.LogError("Error while SendOTP", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return
		}
	}
	token, err := auth.token.CreateSessionToken(value, req.Action, time.Minute*5)
	if err != nil {
		auth.log.LogError("Error while CreateSessionToken", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	w.Header().Set("session-token", token)
	helpers.WriteJSON(w, http.StatusOK, "OTP sent successfully")
}

func (auth *AuthController) verifyOTP(w http.ResponseWriter, payload *token.SessionPayload, email, phone, code, action string) bool {
	if payload.TokenType != action {
		auth.log.LogError("Payload doesnot match")
		helpers.ErrorJson(w, http.StatusForbidden, "Unauthorized: Payload doesnot match")
		return false
	}

	column, value := utils.IdentifiesColumnValue(email, phone)
	if action == utils.SIGN_UP {
		if len(column) == 0 {
			auth.log.LogError("Error while IdentifiesColumnValue", column)
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return false
		}
		res, err := auth.storage.CheckDataExist(column, value)
		if err != nil {
			auth.log.LogError("Error while CheckDataExist", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return false
		}
		if res {
			fmtError := fmt.Errorf("account already exist using this %s", column)
			auth.log.LogError(fmtError)
			helpers.ErrorJson(w, http.StatusConflict, fmtError.Error())
			return false
		}
	}

	if !utils.IsEmpty(email) {
		if payload.Value != email {
			err := errors.New("Forbidden")
			auth.log.LogError("Forbidden", err)
			helpers.ErrorJson(w, http.StatusForbidden, "Forbidden")
			return false
		}
		status, err := auth.emailService.VerfiyOTP(email, code, auth.redis)
		if err != nil {
			auth.log.LogError("Error while VerfiyOTP", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return false
		}
		if !status {
			err = errors.New("Invalid OTP")
			auth.log.LogError("Invalid OTP", err)
			helpers.ErrorJson(w, http.StatusUnauthorized, "Invalid OTP")
			return false
		}
	} else {
		if payload.Value != phone {
			err := errors.New("Forbidden")
			auth.log.LogError("Forbidden", err)
			helpers.ErrorJson(w, http.StatusForbidden, "Forbidden")
			return false
		}
		phoneNumber := fmt.Sprintf("+91%s", phone)
		status, err := auth.twilioService.VerfiyOTP(phoneNumber, code)
		if err != nil {
			auth.log.LogError("Error while VerfiyOTP", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return false
		}
		if !status {
			err = errors.New("Invalid OTP")
			auth.log.LogError("Invalid OTP", err)
			helpers.ErrorJson(w, http.StatusUnauthorized, "Invalid OTP")
			return false
		}
	}
	return true
}

func (auth *AuthController) SignUpUser(w http.ResponseWriter, r *http.Request) {
	req := new(models.SignupReq)

	err := helpers.ValidateBody(r, req)
	if err != nil {
		auth.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, InvalidBody)
		return
	}

	err = utils.ValidateEmailOrPhone(req.Email, req.Phone)
	if err != nil {
		auth.log.LogError("Error while ValidateEmailOrPhone", err)
		helpers.ErrorJson(w, http.StatusBadRequest, "Invalid Email or Phone")
		return
	}

	payload := r.Context().Value(middlewares.AuthorizationPayloadKey).(*token.SessionPayload)
	verify := auth.verifyOTP(w, payload, req.Email, req.Phone, req.Code, utils.SIGN_UP)
	if !verify {
		return
	}

	err = auth.storage.InsertUser(req)
	if err != nil {
		err = fmt.Errorf("error while inserting user %w", err)
		auth.log.LogError("Error while InsertUser", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	token, err := auth.token.CreateAccessToken(req.UserName, time.Minute*5)
	if err != nil {
		auth.log.LogError("Error while CreateAccessToken", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	w.Header().Set("access-token", token)
	helpers.WriteJSON(w, http.StatusOK, "User created Successfully")
}

func (auth *AuthController) LoginUser(w http.ResponseWriter, r *http.Request) {
	req := new(models.LoginReq)

	err := helpers.ValidateBody(r, req)
	if err != nil {
		auth.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, InvalidBody)
		return
	}

	res, err := auth.storage.CheckDataExist("user_name", req.UserName)
	if err != nil {
		auth.log.LogError("Error while CheckDataExist", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	if !res {
		err = fmt.Errorf("user doesnt exist %w", err)
		auth.log.LogError("User doesn't", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, "User doesnt exist")
		return
	}

	res, err = auth.storage.CheckPassword(req.UserName, req.Password)
	if err != nil {
		auth.log.LogError("Error while CheckPassword", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	if !res {
		err = fmt.Errorf("wrong password %w", err)
		auth.log.LogError("Wrong password", err)
		helpers.ErrorJson(w, http.StatusForbidden, "User crediantials doesn't match")
		return
	}

	token, err := auth.token.CreateAccessToken(req.UserName, time.Minute*10)
	if err != nil {
		auth.log.LogError("Error while CreateAccessToken", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	w.Header().Set("access-token", token)
	helpers.WriteJSON(w, http.StatusOK, "User loggedin Successfully")
}

func (auth *AuthController) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	req := new(models.ForgotPassword)

	err := helpers.ValidateBody(r, req)
	if err != nil {
		auth.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, InvalidBody)
		return
	}

	err = utils.ValidateEmailOrPhone(req.Email, req.Phone)
	if err != nil {
		auth.log.LogError("Error while ValidateEmailOrPhone", err)
		helpers.ErrorJson(w, http.StatusBadRequest, "Invalid Email or Phone")
		return
	}

	payload := r.Context().Value(middlewares.AuthorizationPayloadKey).(*token.SessionPayload)
	verify := auth.verifyOTP(w, payload, req.Email, req.Phone, req.Code, utils.FORGOT_PASSWORD)
	if !verify {
		return
	}

	err = auth.storage.ChangePassword(req)
	if err != nil {
		auth.log.LogError("Error while ChangePassword", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "Password changed successfully")

}
