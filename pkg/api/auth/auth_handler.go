package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/akmal4410/gestapo/internal/config"
	"github.com/akmal4410/gestapo/pkg/api/auth/database"
	"github.com/akmal4410/gestapo/pkg/api/auth/database/entity"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/service/cache"
	"github.com/akmal4410/gestapo/pkg/service/logger"
	"github.com/akmal4410/gestapo/pkg/service/mail"
	"github.com/akmal4410/gestapo/pkg/service/sso"
	"github.com/akmal4410/gestapo/pkg/service/token"
	"github.com/akmal4410/gestapo/pkg/service/twilio"
	"github.com/akmal4410/gestapo/pkg/utils"
)

const (
	SsoOAuthString      string = "sso-oauth-string"
	InternalServerError string = "Internal server error"
	InvalidBody         string = "Invalid Body"
)

type AuthHandler struct {
	twilioService twilio.TwilioService
	emailService  mail.EmailService
	storage       *database.AuthStore
	token         token.Maker
	redis         cache.Cache
	log           logger.Logger
	config        *config.Config
}

func NewAuthHandler(
	twilio twilio.TwilioService,
	email mail.EmailService,
	storage *database.AuthStore,
	token token.Maker,
	redisCache cache.Cache,
	logger logger.Logger,
	config *config.Config,
) *AuthHandler {
	return &AuthHandler{
		twilioService: twilio,
		emailService:  email,
		storage:       storage,
		token:         token,
		redis:         redisCache,
		log:           logger,
		config:        config,
	}
}

func (auth *AuthHandler) SendOTP(w http.ResponseWriter, r *http.Request) {
	req := new(entity.SendOTPReq)

	err := helpers.ValidateBody(r, req)
	if err != nil {
		auth.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, InvalidBody)
		return
	}

	err = helpers.ValidateEmailOrPhone(req.Email, req.Phone)
	if err != nil {
		auth.log.LogError("Error while ValidateEmailOrPhone", err)
		helpers.ErrorJson(w, http.StatusBadRequest, "Invalid Email or Phone")
		return
	}

	column, value := helpers.IdentifiesColumnValue(req.Email, req.Phone)
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

	if !helpers.IsEmpty(req.Email) {
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

func (auth *AuthHandler) verifyOTP(w http.ResponseWriter, payload *token.SessionPayload, email, phone, code, action string) bool {
	if payload.TokenType != action {
		auth.log.LogError("Payload doesnot match")
		helpers.ErrorJson(w, http.StatusForbidden, "Unauthorized: Payload doesnot match")
		return false
	}

	column, value := helpers.IdentifiesColumnValue(email, phone)
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

	if !helpers.IsEmpty(email) {
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
			err = errors.New("invalid OTP")
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
			err = errors.New("invalid OTP")
			auth.log.LogError("Invalid OTP", err)
			helpers.ErrorJson(w, http.StatusUnauthorized, "Invalid OTP")
			return false
		}
	}
	return true
}

func (auth *AuthHandler) SignUpUser(w http.ResponseWriter, r *http.Request) {
	req := new(entity.SignupReq)

	err := helpers.ValidateBody(r, req)
	if err != nil {
		auth.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, InvalidBody)
		return
	}

	err = helpers.ValidateEmailOrPhone(req.Email, req.Phone)
	if err != nil {
		auth.log.LogError("Error while ValidateEmailOrPhone", err)
		helpers.ErrorJson(w, http.StatusBadRequest, "Invalid Email or Phone")
		return
	}

	payload := r.Context().Value(AuthorizationPayloadKey).(*token.SessionPayload)
	verify := auth.verifyOTP(w, payload, req.Email, req.Phone, req.Code, utils.SIGN_UP)
	if !verify {
		return
	}

	id, err := auth.storage.InsertUser(req)
	if err != nil {
		err = fmt.Errorf("error while inserting user %w", err)
		auth.log.LogError("Error while InsertUser", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	token, err := auth.token.CreateAccessToken(id, req.UserName, req.UserType, time.Minute*5)
	if err != nil {
		auth.log.LogError("Error while CreateAccessToken", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	w.Header().Set("access-token", token)
	helpers.WriteJSON(w, http.StatusOK, "User Signup Successfully")
}

func (auth *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	req := new(entity.LoginReq)

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

	payload, err := auth.storage.GetTokenPayload("user_name", req.UserName)
	if err != nil {
		auth.log.LogError("Error while GetTokenPayload", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	token, err := auth.token.CreateAccessToken(payload.UserId, req.UserName, payload.UserType, time.Minute*10)
	if err != nil {
		auth.log.LogError("Error while CreateAccessToken", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	w.Header().Set("access-token", token)
	helpers.WriteJSON(w, http.StatusOK, "User loggedin Successfully")
}

func (auth *AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	req := new(entity.ForgotPassword)

	err := helpers.ValidateBody(r, req)
	if err != nil {
		auth.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, InvalidBody)
		return
	}

	err = helpers.ValidateEmailOrPhone(req.Email, req.Phone)
	if err != nil {
		auth.log.LogError("Error while ValidateEmailOrPhone", err)
		helpers.ErrorJson(w, http.StatusBadRequest, "Invalid Email or Phone")
		return
	}

	payload := r.Context().Value(AuthorizationPayloadKey).(*token.SessionPayload)
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

func (auth *AuthHandler) SSOAuth(w http.ResponseWriter, r *http.Request) {
	req := new(entity.SsoReq)

	err := helpers.ValidateBody(r, req)
	if err != nil {
		auth.log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(w, http.StatusBadRequest, InvalidBody)
		return
	}

	token := r.Context().Value(AuthorizationPayloadKey).(string)

	var email, fullname string

	switch req.Action {
	case utils.SSO_ANDROID:
		email, fullname, err = sso.GoogleOauth(token, auth.config.OAuth.AndroidClientId, auth.log)
		if err != nil {
			if err.Error() == "missing claims" {
				helpers.ErrorJson(w, http.StatusConflict, "conflict occurs, missing claims")
				return
			}
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return
		}
	case utils.SSO_IOS:
		email, fullname, err = sso.GoogleOauth(token, auth.config.OAuth.IOSClientId, auth.log)
		if err != nil {
			if err.Error() == "missing claims" {
				helpers.ErrorJson(w, http.StatusConflict, "conflict occurs, missing claims")
				return
			}
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return
		}
	default:
		helpers.ErrorJson(w, http.StatusBadRequest, "Invalid action")
		return
	}

	//checks if the user exist or not
	exist, err := auth.storage.CheckDataExist("email", email)
	if err != nil {
		auth.log.LogError("Error while CheckDataExist", err)
		helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	//already exist so login
	if exist {
		payload, err := auth.storage.GetTokenPayload("email", email)
		if err != nil {
			auth.log.LogError("Error while GetTokenPayload", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return
		}

		token, err := auth.token.CreateAccessToken(payload.UserId, payload.UserName, payload.UserType, time.Minute*10)
		if err != nil {
			auth.log.LogError("Error while CreateAccessToken", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return
		}

		w.Header().Set("access-token", token)
		helpers.WriteJSON(w, http.StatusOK, "User loggedin Successfully")
	} else {
		signupReq := entity.SignupReq{
			Email:    email,
			UserName: fullname,
			UserType: req.UserType,
			Password: email + fullname + req.UserType,
		}
		id, err := auth.storage.InsertUser(&signupReq)
		if err != nil {
			auth.log.LogError("Error while InsertUser", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return
		}

		token, err := auth.token.CreateAccessToken(id, fullname, req.UserType, time.Minute*5)
		if err != nil {
			auth.log.LogError("Error while CreateAccessToken", err)
			helpers.ErrorJson(w, http.StatusInternalServerError, InternalServerError)
			return
		}
		w.Header().Set("access-token", token)
		helpers.WriteJSON(w, http.StatusOK, "User Signup Successfully")
	}

}
