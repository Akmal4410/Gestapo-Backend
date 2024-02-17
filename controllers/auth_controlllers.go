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
	"github.com/akmal4410/gestapo/services/mail"
	"github.com/akmal4410/gestapo/services/password"
	"github.com/akmal4410/gestapo/services/token"
	"github.com/akmal4410/gestapo/services/twilio"
	"github.com/akmal4410/gestapo/utils"
	"github.com/pkg/errors"
)

type AuthController struct {
	twilioService twilio.TwilioService
	emailService  mail.EmailService
	storage       *database.Storage
	token         token.Maker
	redis         cache.Cache
}

func NewAuthController(
	twilio twilio.TwilioService,
	email mail.EmailService,
	storage *database.Storage,
	token token.Maker,
	redisCache cache.Cache,
) *AuthController {
	return &AuthController{
		twilioService: twilio,
		emailService:  email,
		storage:       storage,
		token:         token,
		redis:         redisCache,
	}
}

func (auth AuthController) SendOTP(w http.ResponseWriter, r *http.Request) {
	req := new(models.SendOTPReq)

	err := helpers.ValidateBody(r, req)
	if err != nil {
		helpers.ErrorJson(w, http.StatusBadRequest, err)
		return
	}

	err = utils.ValidateEmailOrPhone(req.Email, req.Phone)
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
		err = auth.emailService.SendEmail(req.Email, utils.EmailSubject, utils.EmailSubject, auth.redis)
		if err != nil {
			helpers.ErrorJson(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		phoneNumber := fmt.Sprintf("+91%s", req.Phone)
		err = auth.twilioService.SendOTP(phoneNumber)
		if err != nil {
			helpers.ErrorJson(w, http.StatusInternalServerError, err)
			return
		}
	}
	token, err := auth.token.CreateSessionToken(value, time.Minute*5)
	if err != nil {
		helpers.ErrorJson(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("session-token", token)
	helpers.WriteJSON(w, http.StatusOK, "OTP sent successfully")
}

func (auth AuthController) SignupUser(w http.ResponseWriter, r *http.Request) {
	req := new(models.SignupReq)
	err := helpers.ValidateBody(r, req)
	if err != nil {
		helpers.ErrorJson(w, http.StatusBadRequest, errors.New("invalid Body"))
		return
	}

	err = utils.ValidateEmailOrPhone(req.Email, req.Phone)
	if err != nil {
		helpers.ErrorJson(w, http.StatusBadRequest, fmt.Errorf("invalid %w", err))
		return
	}

	column, value := utils.IdentifiesColumnValue(req.Email, req.Phone)
	if len(column) == 0 {
		helpers.ErrorJson(w, http.StatusBadRequest, fmt.Errorf("invalid %w", err))
		return
	}
	res, err := auth.storage.CheckUserExist(column, value)
	if err != nil {
		helpers.ErrorJson(w, http.StatusInternalServerError, errors.New("user already exist"))
		return
	}
	if res {
		helpers.ErrorJson(w, http.StatusInternalServerError, fmt.Errorf("account already exist using this %s", column))
		return
	}
	payload := r.Context().Value(middlewares.AuthorizationPayloadKey).(*token.SessionPayload)
	if !utils.IsEmpty(req.Email) {
		if payload.Value != req.Email {
			err = errors.New("Forbidden")
			helpers.ErrorJson(w, http.StatusForbidden, err)
			return
		}
		status, err := auth.emailService.VerfiyOTP(req.Email, req.Code, auth.redis)
		if err != nil {
			helpers.ErrorJson(w, http.StatusInternalServerError, err)
			return
		}
		if !status {
			err = errors.New("Invalid otp")
			helpers.ErrorJson(w, http.StatusUnauthorized, err)
			return
		}
	} else {
		if payload.Value != req.Phone {
			err = errors.New("Forbidden")
			helpers.ErrorJson(w, http.StatusForbidden, err)
			return
		}
		phoneNumber := fmt.Sprintf("+91%s", req.Phone)
		status, err := auth.twilioService.VerfiyOTP(phoneNumber, req.Code)
		if err != nil {
			helpers.ErrorJson(w, http.StatusInternalServerError, err)
			return
		}
		if !status {
			err = errors.New("Invalid otp")
			helpers.ErrorJson(w, http.StatusUnauthorized, err)
			return
		}
	}

	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		helpers.ErrorJson(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(hashedPassword)

	token, err := auth.token.CreateAccessToken(value, time.Minute*5)
	if err != nil {
		helpers.ErrorJson(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("access-token", token)
	helpers.WriteJSON(w, http.StatusOK, "OTP sent successfully")

}

func (auth AuthController) LoginUser(w http.ResponseWriter, r *http.Request) {
}
