package models

import (
	"time"

	"gorm.io/gorm"
)

type User_Data struct {
	ID            uint `gorm:"NOT NULL;PRIMARY_KEY;AUTO_INCREMENT"`
	Profile_Image string
	Full_Name     string
	User_Name     string `gorm:"NOT NULL;UNIQUE"`
	Phone         string `gorm:"NOT NULL;UNIQUE"`
	DOB           time.Time
	Gender        string
	User_type     string `gorm:"NOT NULL;CHECK:user_type = 'USER' OR user_type = 'Merchant' OR user_type = 'ADMIN'"`
	Password      string `gorm:"NOT NULL"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
