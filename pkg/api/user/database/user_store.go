package database

import "github.com/akmal4410/gestapo/internal/database"

type UserStore struct {
	storage *database.Storage
}

func NewUserStore(storage *database.Storage) *UserStore {
	return &UserStore{
		storage: storage,
	}
}

func (store *UserStore) GetDiscount() error {
	return nil
}
