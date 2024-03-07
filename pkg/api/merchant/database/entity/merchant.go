package entity

import (
	"time"
)

type GetMerchantRes struct {
	ID           string     `json:"id"`
	ProfileImage *string    `json:"profile_image,omitempty"`
	FullName     *string    `json:"full_name,omitempty"`
	UserName     string     `json:"user_name,omitempty"`
	Phone        *string    `json:"phone,omitempty"`
	Email        *string    `json:"email,omitempty"`
	DOB          *time.Time `json:"dob,omitempty"`
	Gender       *string    `json:"gender,omitempty"`
	UserType     string     `json:"user_type,omitempty"`
}

type EditMerchantReq struct {
	ProfileImage string `json:"profile_image" validate:"omitempty"`
	FullName     string `json:"full_name" validate:"omitempty"`
	DOB          string `json:"dob" validate:"omitempty,validate_date"`
	Gender       string `json:"gender" validate:"omitempty,gender"`
}
