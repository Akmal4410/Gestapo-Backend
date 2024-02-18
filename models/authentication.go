package models

type SendOTPReq struct {
	Email string `json:"email" validate:"omitempty"`
	Phone string `json:"phone" validate:"omitempty,len=10,numeric"`
}

type SignupReq struct {
	Email    string `json:"email" validate:"omitempty"`
	Phone    string `json:"phone" validate:"omitempty,len=10,numeric"`
	UserName string `json:"user_name" validate:"required,min=4,max=12"`
	UserType string `json:"user_type" validate:"user_type"`
	Code     string `json:"code" validate:"required,min=6,max=6"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

// type SignupReq struct {
// 	Email        string    `json:"email" validate:"omitempty"`
// 	Phone        string    `json:"phone" validate:"omitempty,len=10,numeric"`
// 	Code         string    `json:"code" validate:"required"`
// 	ProfileImage string    `json:"profile_image"`
// 	FullName     string    `json:"full_name" validate:"required"`
// 	UserName     string    `json:"user_name" validate:"required"`
// 	DOB          time.Time `json:"dob"`
// 	Gender       string    `json:"gender" validate:"required"`
// 	User_type    string    `json:"user_type" validate:"UserType"`
// 	Password     string    `json:"password" validate:"required"`
// }
