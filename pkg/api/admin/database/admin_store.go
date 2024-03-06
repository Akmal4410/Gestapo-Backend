package database

import (
	"time"

	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/api/admin/database/entity"
	"github.com/google/uuid"
)

type AdminStore struct {
	storage *database.Storage
}

func NewAdminStore(storage *database.Storage) *AdminStore {
	return &AdminStore{storage: storage}
}

func (store *AdminStore) CheckCategoryExist(category string) (bool, error) {
	checkQuery := `SELECT * FROM categories WHERE category_name = $1;`
	res, err := store.storage.DB.Exec(checkQuery, category)
	if err != nil {
		return false, err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return result != 0, nil
}

func (store AdminStore) InsertCategory(req *entity.InsertCategoryReq) error {
	createdAt := time.Now()
	updatedAt := time.Now()

	insertQuery := `
	INSERT INTO categories (id, category_name, created_at, updated_at)
	VALUES ($1, $2, $3, $4);
	`

	uuId, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	_, err = store.storage.DB.Exec(insertQuery, uuId, req.Category_Name, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}
