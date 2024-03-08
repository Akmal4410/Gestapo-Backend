package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/akmal4410/gestapo/pkg/utils"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateBody(body io.Reader, data any) error {
	RegisterValidator()
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return err
	}
	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}

func RegisterValidator() {
	err := validate.RegisterValidation("user_type", validateUserType)
	if err != nil {
		fmt.Println("Error registering user_type:", err.Error())
	}
	err = validate.RegisterValidation("signup_action", validateSignupAction)
	if err != nil {
		fmt.Println("Error registering signup_action:", err.Error())
	}
	err = validate.RegisterValidation("sso_action", validateSSOAction)
	if err != nil {
		fmt.Println("Error registering sso_action:", err.Error())
	}
	err = validate.RegisterValidation("gender", validateGender)
	if err != nil {
		fmt.Println("Error registering gender:", err.Error())
	}
	err = validate.RegisterValidation("validate_date", validateDate)
	if err != nil {
		fmt.Println("Error registering validate_date:", err.Error())
	}
}

var validateUserType validator.Func = func(fl validator.FieldLevel) bool {
	if userType, ok := fl.Field().Interface().(string); ok {
		// Check userType is supported or not
		return utils.IsSupportedUsers(userType)
	}
	return false
}

var validateSignupAction validator.Func = func(fl validator.FieldLevel) bool {
	if signupAction, ok := fl.Field().Interface().(string); ok {
		// Check signupAction is supported or not
		return utils.IsSupportedSignupAction(signupAction)
	}
	return false
}

var validateSSOAction validator.Func = func(fl validator.FieldLevel) bool {
	if signupAction, ok := fl.Field().Interface().(string); ok {
		// Check sso-action is supported or not
		return utils.IsSupportedSSOAction(signupAction)
	}
	return false
}

var validateGender validator.Func = func(fl validator.FieldLevel) bool {
	if signupAction, ok := fl.Field().Interface().(string); ok {
		// Check gender is supported or not
		return utils.IsSupportedGender(signupAction)
	}
	return false
}

var validateDate validator.Func = func(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
