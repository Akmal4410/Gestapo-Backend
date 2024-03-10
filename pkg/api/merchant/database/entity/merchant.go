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

type AddProductReq struct {
	ProductName   string    `json:"product_name" validate:"required"`
	Description   string    `json:"description" validate:"required"`
	ProductImages []string  `json:"product_images" validate:"omitempty"`
	Sizes         []float64 `json:"sizes" validate:"required"`
	Price         float64   `json:"price" validate:"required"`
	CategoryId    string    `json:"category_id" validate:"required"`
	Quantity      int       `json:"quantity" validate:"required"`
	DiscountName  string    `json:"discount_name" validate:"omitempty"`
	Percent       int32     `json:"percent" validate:"omitempty"`
}
