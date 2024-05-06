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

type Products struct {
	ID          uuid.UUID  `gorm:"NOT NULL;PRIMARY_KEY"`
	MerchentID  uuid.UUID  `gorm:"NOT NULL"`
	Category    Categories `gorm:"foreignKey:CategoryID;references:ID"`
	CategoryID  uuid.UUID  `gorm:"NOT NULL;index"`
	Discount    Discounts  `gorm:"foreignKey:DiscountID;references:ID"`
	DiscountID  *uuid.UUID
	ProductName string          `gorm:"NOT NULL"`
	Description string          `gorm:"NOT NULL"`
	Images      pq.StringArray  `gorm:"type:text[]"`
	Size        pq.Float64Array `gorm:"type:float[]"`
	Price       float64         `gorm:"NOT NULL"`
	CreatedAt   time.Time       `gorm:"NOT NULL"`
	UpdatedAt   time.Time       `gorm:"NOT NULL"`
	DeletedAt   gorm.DeletedAt
}

type Discounts struct {
	ID          uuid.UUID `gorm:"NOT NULL; PRIMARY_KEY"`
	MerchentID  uuid.UUID `gorm:"NOT NULL"`
	Name        string    `gorm:"NOT NULL"`
	Description string    `gorm:"NOT NULL"`
	Percent     float64   `gorm:"NOT NULL"`
	CardColor   string    `gorm:"NOT NULL; default:0xFF808080"`
	StartTime   time.Time `gorm:"NOT NULL"`
	EndTime     time.Time `gorm:"NOT NULL"`
	CreatedAt   time.Time `gorm:"NOT NULL"`
	UpdatedAt   time.Time `gorm:"NOT NULL"`
}
type Inventories struct {
	ID        uuid.UUID `gorm:"NOT NULL;PRIMARY_KEY"`
	Product   Products  `gorm:"foreignKey:ProductID;references:ID"`
	ProductID uuid.UUID `gorm:"NOT NULL;index"`
	Size      float64   `gorm:"NOT NULL"`
	Quantity  int       `gorm:"NOT NULL"`
	CreatedAt time.Time `gorm:"NOT NULL"`
	UpdatedAt time.Time `gorm:"NOT NULL"`
}

type Wishlists struct {
	ID        uuid.UUID `gorm:"NOT NULL;PRIMARY_KEY"`
	User      User_Data `gorm:"foreignKey:UserID;references:ID"`
	UserID    uuid.UUID `gorm:"NOT NULL"`
	Product   Products  `gorm:"foreignKey:ProductID;references:ID"`
	ProductID uuid.UUID `gorm:"NOT NULL;index"`
	CreatedAt time.Time `gorm:"NOT NULL"`
	UpdatedAt time.Time `gorm:"NOT NULL"`
}

type Carts struct {
	ID        uuid.UUID `gorm:"NOT NULL;PRIMARY_KEY"`
	User      User_Data `gorm:"foreignKey:UserID;references:ID"`
	UserID    uuid.UUID `gorm:"NOT NULL"`
	Price     float64   `gorm:"NOT NULL"`
	CreatedAt time.Time `gorm:"NOT NULL"`
	UpdatedAt time.Time `gorm:"NOT NULL"`
}

type Cart_Items struct {
	ID          uuid.UUID   `gorm:"NOT NULL;PRIMARY_KEY"`
	Cart        Carts       `gorm:"foreignKey:CartID;references:ID"`
	CartID      uuid.UUID   `gorm:"NOT NULL"`
	Product     Products    `gorm:"foreignKey:ProductID;references:ID"`
	ProductID   uuid.UUID   `gorm:"NOT NULL;index"`
	Inventory   Inventories `gorm:"foreignKey:InventoryID;references:ID"`
	InventoryID uuid.UUID   `gorm:"NOT NULL;index"`
	Quantity    int         `gorm:"NOT NULL"`
	Price       float64     `gorm:"NOT NULL"`
	CreatedAt   time.Time   `gorm:"NOT NULL"`
	UpdatedAt   time.Time   `gorm:"NOT NULL"`
}

type Addresses struct {
	ID          uuid.UUID `gorm:"NOT NULL;PRIMARY_KEY"`
	User        User_Data `gorm:"foreignKey:UserID;references:ID"`
	UserID      uuid.UUID `gorm:"NOT NULL"`
	Title       string    `gorm:"NOT NULL"`
	AddressLine string    `gorm:"NOT NULL"`
	Country     string
	City        string
	PostalCode  int64
	Landmark    string
	IsDefault   bool
	CreatedAt   time.Time `gorm:"NOT NULL"`
	UpdatedAt   time.Time `gorm:"NOT NULL"`
}
