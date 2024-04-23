package database

import (
	"fmt"

	"github.com/akmal4410/gestapo/internal/models"
	"gorm.io/gorm"
)

type DBMigration struct {
	user_data  models.User_Data
	categories models.Categories
	products   models.Products
	carts      models.Carts
}

var migrate DBMigration

func AutoMigrateTables(gormDB *gorm.DB) {
	if err := gormDB.AutoMigrate(&migrate.user_data); err != nil {
		fmt.Println(err.Error())
	}

	if err := gormDB.AutoMigrate(
		&migrate.categories,
		&migrate.products,
	); err != nil {
		fmt.Println(err.Error())
	}
	if err := gormDB.AutoMigrate(&migrate.carts); err != nil {
		fmt.Println(err.Error())
	}

	// Add foreign key constraints
	// if err := gormDB.Model(&migrate.products).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT"); err != nil {
	// 	fmt.Println(err.Error())
	// }
}
