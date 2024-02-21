package models

type SendOTPReq struct {
	Email  string `json:"email" validate:"omitempty"`
	Phone  string `json:"phone" validate:"omitempty,len=10,numeric"`
	Action string `json:"action" validate:"required"`
}

type SignupReq struct {
	Email    string `json:"email" validate:"omitempty"`
	Phone    string `json:"phone" validate:"omitempty,len=10,numeric"`
	UserName string `json:"user_name" validate:"required,min=4,max=12"`
	UserType string `json:"user_type" validate:"USER_TYPE"`
	Code     string `json:"code" validate:"required,min=6,max=6"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}
type LoginReq struct {
	UserName string `json:"user_name" validate:"required,min=4,max=12"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type ForgotPassword struct {
	Email    string `json:"email" validate:"omitempty"`
	Phone    string `json:"phone" validate:"omitempty,len=10,numeric"`
	Code     string `json:"code" validate:"required,min=6,max=6"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}
