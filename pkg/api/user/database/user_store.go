package database

import (
	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/api/user/database/entity"
)

type UserStore struct {
	storage *database.Storage
}

func NewUserStore(storage *database.Storage) *UserStore {
	return &UserStore{
		storage: storage,
	}
}

func (store *UserStore) GetDiscount() (*entity.DiscountRes, error) {
	selectQuery := `
	SELECT 
    p.id AS product_id,
    d.name AS name,
	d.description AS description,
    d.percent AS percent,
    p.images[1] AS image
	FROM products p
	JOIN discounts d ON p.discount_id = d.id
	WHERE d.end_time > NOW()
	ORDER BY d.percent DESC
	LIMIT 1;
	`

	rows := store.storage.DB.QueryRow(selectQuery)
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	var discount entity.DiscountRes
	err := rows.Scan(
		&discount.ProductID,
		&discount.Name,
		&discount.Description,
		&discount.Percentage,
		&discount.ProductImage,
	)
	if err != nil {
		return nil, err
	}
	return &discount, nil
}
