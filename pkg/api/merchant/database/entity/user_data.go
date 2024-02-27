package entity

import (
	"time"
)

type UserData struct {
	ID            string     `json:"id"`
	Profile_Image *string    `json:"profile_image,omitempty"`
	Full_Name     *string    `json:"full_name,omitempty"`
	User_Name     string     `json:"user_name,omitempty"`
	Phone         *string    `json:"phone,omitempty"`
	Email         *string    `json:"email,omitempty"`
	DOB           *time.Time `json:"dob,omitempty"`
	Gender        *string    `json:"gender,omitempty"`
	User_type     string     `json:"user_type,omitempty"`
}
