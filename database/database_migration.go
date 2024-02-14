package database

import (
	"fmt"

	"github.com/akmal4410/gestapo/models"
	"gorm.io/gorm"
)

type DBMigration struct {
	user_data models.User_Data
}

var migrate DBMigration

func AutoMigrateTables(gormDB *gorm.DB) {
	if err := gormDB.AutoMigrate(&migrate.user_data); err != nil {
		fmt.Println(err.Error())
	}
}
