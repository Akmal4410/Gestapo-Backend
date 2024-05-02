package db

import (
	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/grpc_api/product_service/db/entity"
	"github.com/lib/pq"
)

type ProductStore struct {
	storage *database.Storage
}

func NewProductStore(storage *database.Storage) *ProductStore {
	return &ProductStore{storage: storage}

}

func (store *ProductStore) GetProducts() ([]entity.GetProductRes, error) {
	var products []entity.GetProductRes
	selectQuery := `
		SELECT id, product_name, images, price
		FROM products;
	`

	rows, err := store.storage.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.GetProductRes
		var images pq.StringArray

		err := rows.Scan(
			&product.ID,
			&product.ProductName,
			&images,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}
		product.ProductImages = []string(images)
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return []entity.GetProductRes{}, nil
	}

	return products, nil
}

func (store *ProductStore) GetProductById(productId string) (*entity.GetProductRes, error) {
	selectQuery := `
	SELECT
    p.id AS id,
	p.merchent_id AS merchent_id,
    p.product_name AS product_name,
    p.description AS description,
    c.category_name AS category_name,
    p.size AS size,
    p.price AS price,
    CASE
        WHEN d.end_time IS NOT NULL AND d.end_time > NOW()
		THEN p.price - (p.price * d.percent / 100) 
        ELSE NULL
    END AS discount_price,
    p.images AS product_images
	FROM
    products p
	LEFT JOIN
    categories c ON p.category_id = c.id
	LEFT JOIN
    discounts d ON p.discount_id = d.id
	WHERE 
	p.id = $1;
	`
	rows := store.storage.DB.QueryRow(selectQuery, productId)
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var product entity.GetProductRes

	var images pq.StringArray
	var sizes pq.Float64Array

	err := rows.Scan(
		&product.ID,
		&product.MerchantID,
		&product.ProductName,
		&product.Description,
		&product.CategoryName,
		&sizes,
		&product.Price,
		&product.DiscountPrice,
		&images,
	)
	product.ProductImages = []string(images)
	// Convert pq.Float64Array to []float64
	var sizeList []float64
	for _, v := range sizes {
		sizeList = append(sizeList, float64(v))
	}
	product.Size = &sizeList

	if err != nil {
		return nil, err
	}
	return &product, nil
}
