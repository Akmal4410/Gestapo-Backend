package database

import (
	"database/sql"

	"github.com/akmal4410/gestapo/util"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(config util.Config) (*Storage, error) {
	db, err := sql.Open(config.DBServer, config.DBSource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}
