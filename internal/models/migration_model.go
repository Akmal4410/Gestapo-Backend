package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User_Data struct {
	ID            uuid.UUID `gorm:"NOT NULL;PRIMARY_KEY"`
	Profile_Image string
	Full_Name     string
	User_Name     string `gorm:"NOT NULL;UNIQUE"`
	Phone         string `gorm:"UNIQUE;DEFAULT:NULL"`
	Email         string `gorm:"UNIQUE;DEFAULT:NULL"`
	DOB           time.Time
	Gender        string
	User_type     string `gorm:"NOT NULL;CHECK:user_type = 'USER' OR user_type = 'MERCHANT' OR user_type = 'ADMIN'"`
	Password      string `gorm:"NOT NULL"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
