package database

import (
	"database/sql"
	"log"

	"github.com/akmal4410/gestapo/utils"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(config utils.Config) (*Storage, error) {
	db, err := sql.Open(config.DBServer, config.DBSource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(postgres.Open(config.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal("Error executing gorm  - ", err)
		return nil, err
	}

	AutoMigrateTables(gormDB)
	return &Storage{db: db}, nil
}
