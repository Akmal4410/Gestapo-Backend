package helpers

import (
	"fmt"

	"github.com/akmal4410/gestapo/utils"
	"github.com/go-playground/validator/v10"
)

var validateUserType validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if userType, ok := fieldLevel.Field().Interface().(string); ok {
		//check userType is supported or not
		return utils.IsSupportedUsers(userType)
	}
	return false
}

func RegisterValidator() {
	validate := validator.New()
	err := validate.RegisterValidation("UserType", validateUserType)
	if err != nil {
		fmt.Println("Error registering custom validation :", err.Error())
	}

}
