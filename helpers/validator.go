package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akmal4410/gestapo/utils"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateBody(r *http.Request, data any) error {
	RegisterValidator()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return err
	}
	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}

func RegisterValidator() {
	err := validate.RegisterValidation("USER_TYPE", validateUserType)
	if err != nil {
		fmt.Println("Error registering custom validation:", err.Error())
	}
	err = validate.RegisterValidation("SIGNUP_ACTION", validateSignupAction)
	if err != nil {
		fmt.Println("Error registering custom validation:", err.Error())
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
