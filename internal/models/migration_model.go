package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
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
	User_type     string    `gorm:"NOT NULL;CHECK:user_type = 'USER' OR user_type = 'MERCHANT' OR user_type = 'ADMIN'"`
	Password      string    `gorm:"NOT NULL"`
	CreatedAt     time.Time `gorm:"NOT NULL"`
	UpdatedAt     time.Time `gorm:"NOT NULL"`
	DeletedAt     gorm.DeletedAt
}

//PRODUCTS------------------------------

type Categories struct {
	ID            uuid.UUID `gorm:"NOT NULL;PRIMARY_KEY"`
	Category_Name string    `gorm:"NOT NULL;UNIQUE"`
	CreatedAt     time.Time `gorm:"NOT NULL"`
	UpdatedAt     time.Time `gorm:"NOT NULL"`
	DeletedAt     gorm.DeletedAt
}

type Inventories struct {
	ID        uuid.UUID `gorm:"NOT NULL;PRIMARY_KEY"`
	Quantity  int       `gorm:"NOT NULL"`
	CreatedAt time.Time `gorm:"NOT NULL"`
	UpdatedAt time.Time `gorm:"NOT NULL"`
	DeletedAt gorm.DeletedAt
}

type Discounts struct {
	ID        uuid.UUID `gorm:"NOT NULL; PRIMARY_KEY"`
	Name      string    `gorm:"NOT NULL;UNIQUE"`
	Percent   int32     `gorm:"NOT NULL"`
	CreatedAt time.Time `gorm:"NOT NULL"`
	UpdatedAt time.Time `gorm:"NOT NULL"`
	DeletedAt gorm.DeletedAt
}

type Products struct {
	ID          uuid.UUID       `gorm:"NOT NULL;PRIMARY_KEY"`
	Category    Categories      `gorm:"foreignKey:CategoryID;references:ID"`
	CategoryID  uuid.UUID       `gorm:"NOT NULL;index"`
	Inventory   Inventories     `gorm:"foreignKey:InventoryID;references:ID"`
	InventoryID uuid.UUID       `gorm:"NOT NULL;index"`
	Discount    Discounts       `gorm:"foreignKey:DiscountID;references:ID"`
	DiscountID  uuid.UUID       `gorm:"NOT NULL;index"`
	ProductName string          `gorm:"NOT NULL"`
	Description string          `gorm:"NOT NULL"`
	Images      pq.StringArray  `gorm:"type:text[]"`
	Size        pq.Float32Array `gorm:"type:float[]"`
	Price       float64         `gorm:"NOT NULL"`
	CreatedAt   time.Time       `gorm:"NOT NULL"`
	UpdatedAt   time.Time       `gorm:"NOT NULL"`
	DeletedAt   gorm.DeletedAt
}
