package models

type OTPData struct {
	PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
}

type VerfiyOTP struct {
	PhoneNumber *OTPData `json:"phone_number,omitempty" validate:"required"`
	Code        string   `json:"code,omitempty" validate:"required"`
}
