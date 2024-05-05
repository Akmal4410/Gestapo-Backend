package database

import (
	"fmt"

	"github.com/akmal4410/gestapo/internal/models"
	"gorm.io/gorm"
)

type DBMigration struct {
	user_data   models.User_Data
	categories  models.Categories
	products    models.Products
	inventories models.Inventories
	dicounts    models.Discounts
	wishlists   models.Wishlists
	carts       models.Carts
	carts_tems  models.Cart_Items
}

var migrate DBMigration

func AutoMigrateTables(gormDB *gorm.DB) {
	if err := gormDB.AutoMigrate(&migrate.user_data); err != nil {
		fmt.Println(err.Error())
	}

	if err := gormDB.AutoMigrate(&migrate.categories); err != nil {
		fmt.Println(err.Error())
	}

	if err := gormDB.AutoMigrate(&migrate.products); err != nil {
		fmt.Println(err.Error())
	}

	if err := gormDB.AutoMigrate(&migrate.inventories); err != nil {
		fmt.Println(err.Error())
	}

	if err := gormDB.AutoMigrate(&migrate.dicounts); err != nil {
		fmt.Println(err.Error())
	}

	if err := gormDB.AutoMigrate(&migrate.wishlists); err != nil {
		fmt.Println(err.Error())
	}

	if err := gormDB.AutoMigrate(&migrate.carts); err != nil {
		fmt.Println(err.Error())
	}

	if err := gormDB.AutoMigrate(&migrate.carts_tems); err != nil {
		fmt.Println(err.Error())
	}
}
