package utils

import (
	"errors"
	"regexp"
	"strings"

	"github.com/akmal4410/gestapo/models"
)

func IsEmpty(str string) bool {
	str = strings.TrimSpace(str)
	if str == "" || len(str) <= 0 {
		return true
	}
	return false
}

func IdentifiesColumnValue(email, phone string) (string, string) {
	// Regular expression patterns for email and phone number validation
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phonePattern := `^\+?[1-9]\d{1,14}$`

	// Check if the input matches the email pattern
	if matched, _ := regexp.MatchString(emailPattern, email); matched {
		return "email", email
	}

	// Check if the input matches the phone number pattern
	if matched, _ := regexp.MatchString(phonePattern, phone); matched {
		return "phone", phone
	}
	return "", ""
}

func EmailOrPhone(input string) string {
	// Regular expression patterns for email and phone number validation
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phonePattern := `^\+?[1-9]\d{1,14}$`

	// Check if the input matches the email pattern
	if matched, _ := regexp.MatchString(emailPattern, input); matched {
		return "email"
	}

	// Check if the input matches the phone number pattern
	if matched, _ := regexp.MatchString(phonePattern, input); matched {
		return "phone"
	}
	return "unknown"
}

func ValidateEmailOrPhone(req *models.SendOTPReq) error {
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
		phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
		if !phoneRegex.MatchString(req.Email) {
			return errors.New("phone number should be 10 digits")
		}

	}
	return nil
}
