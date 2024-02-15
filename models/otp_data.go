package models

type SendOTPReq struct {
	Phone string `json:"phone" validate:"required"`
}

type VerfiyOTPReq struct {
	Phone string `json:"phone" validate:"required"`
	Code  string `json:"code" validate:"required"`
}
