package db

import (
	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/grpc_api/user_service/db/entity"
	"github.com/akmal4410/gestapo/pkg/utils"
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

func (store *UserStore) GetMerchants() ([]entity.MerchantRes, error) {
	selectQuery := `
	SELECT id, full_name, profile_image 
	FROM user_data
	WHERE user_type = $1
	LIMIT 7;
	`
	rows, err := store.storage.DB.Query(selectQuery, utils.MERCHANT)
	if err != nil {
		return nil, err
	}
	var merchants []entity.MerchantRes
	defer rows.Close()
	for rows.Next() {
		var merchant entity.MerchantRes

		err := rows.Scan(
			&merchant.MerchantID,
			&merchant.Name,
			&merchant.ImageURL,
		)
		if err != nil {
			return nil, err
		}
		merchants = append(merchants, merchant)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return merchants, nil
}
