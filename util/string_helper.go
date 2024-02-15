package util

import "regexp"

func IsEmailOrPhone(input string) string {
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
