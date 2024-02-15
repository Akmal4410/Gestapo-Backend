package models

type SendOTPReq struct {
	Email string `json:"email" validate:"omitempty"`
	Phone string `json:"phone" validate:"omitempty,len=10,numeric"`
}

type VerfiyOTPReq struct {
	Email string `json:"email" validate:"omitempty"`
	Phone string `json:"phone" validate:"omitempty,len=10,numeric"`
	Code  string `json:"code" validate:"required"`
}
