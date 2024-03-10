package entity

import "time"

type AddCategoryReq struct {
	Category_Name string `json:"category_name" validate:"required"`
}

type GetCategoriesRes struct {
	ID       string `json:"id"`
	Category string `json:"category"`
}

type GetUserRes struct {
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
