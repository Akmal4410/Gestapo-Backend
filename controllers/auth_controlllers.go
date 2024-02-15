package controllers

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/akmal4410/gestapo/helpers"
	"github.com/akmal4410/gestapo/models"
	"github.com/akmal4410/gestapo/services/mail"
	"github.com/akmal4410/gestapo/services/twilio"
)

type AuthController struct {
	twilioService twilio.TwilioService
	emailService  mail.EmailService
}

func NewAuthController(twilio twilio.TwilioService, email mail.EmailService) *AuthController {
	return &AuthController{
		twilioService: twilio,
		emailService:  email,
	}
}

func (auth AuthController) SendOTP(w http.ResponseWriter, r *http.Request) {
	payload := new(models.SendOTPReq)

	err := helpers.ValidateBody(r, payload)
	if err != nil {
		helpers.ErrorJson(w, http.StatusBadRequest, err)
		return
	}
	err = validate(payload)
	if err != nil {
		helpers.ErrorJson(w, http.StatusBadRequest, err)
		return
	}

	err = auth.emailService.SendEmail(payload.Email, "Sign Up OTP", "Welcome to Gestapo !!!. Use the following OTP to complete your Sign Up procedures. OTP is valid for 5minutes", nil, nil, nil)
	if err != nil {
		helpers.ErrorJson(w, http.StatusInternalServerError, err)
		return
	}
	// phoneNumber := fmt.Sprintf("+91%s", payload.Phone)
	// err = auth.twilioService.SendOTP(phoneNumber)
	// if err != nil {
	// 	helpers.ErrorJson(w, http.StatusInternalServerError, err)
	// 	return
	// }
	helpers.WriteJSON(w, http.StatusOK, "")
}

func (auth AuthController) LoginUser(w http.ResponseWriter, r *http.Request) {
}

func (auth AuthController) CreateUser(w http.ResponseWriter, r *http.Request) {
}

func validate(req *models.SendOTPReq) error {
	// Check that either Email or Phone is present, but not both
	if (req.Email != "" && req.Phone != "") || (req.Email == "" && req.Phone == "") {
		return errors.New("either Email or Phone should be present")
	}
	// Check that at least one field is non-empty
	if req.Email == "" && req.Phone == "" {
		return errors.New("at least one of Email or Phone should be non-empty")
	}
	// Validate email format
	if req.Email != "" {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(req.Email) {
			return errors.New("invalid email format")
		}
	}
	// Validate phone number length
	if req.Phone != "" && len(req.Phone) != 10 {
		return errors.New("phone number should be 10 digits")
	}

	return nil
}
