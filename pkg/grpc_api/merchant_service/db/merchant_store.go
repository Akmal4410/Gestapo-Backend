package db

import (
	"context"
	"fmt"
	"time"

	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/grpc_api/merchant_service/db/entity"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type MarchantStore struct {
	storage *database.Storage
}

func NewMarchantStore(storage *database.Storage) *MarchantStore {
	return &MarchantStore{storage: storage}

}

func (store MarchantStore) CheckDataExist(table, column, value string) (bool, error) {
	checkQuery := fmt.Sprintf(`SELECT * FROM %s WHERE %s = $1;`, table, column)
	res, err := store.storage.DB.Exec(checkQuery, value)
	if err != nil {
		return false, err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return result != 0, nil
}

func (store *MarchantStore) GetProfile(userId string) (*entity.GetMerchantRes, error) {
	selectQuery := `
	SELECT id, profile_image, full_name, user_name, phone, email, dob, gender, user_type 
	FROM user_data WHERE id = $1;
	`
	rows := store.storage.DB.QueryRow(selectQuery, userId)
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	var user entity.GetMerchantRes
	err := rows.Scan(
		&user.ID,
		&user.ProfileImage,
		&user.FullName,
		&user.UserName,
		&user.Phone,
		&user.Email,
		&user.DOB,
		&user.Gender,
		&user.UserType,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (store *MarchantStore) UpdateProfile(id string, req *entity.EditMerchantReq) error {
	updatedAt := time.Now()

	updateQuery := `UPDATE user_data
	SET profile_image = $2, full_name = $3, dob = $4, gender = $5, updated_at = $6
	WHERE id = $1;`

	res, err := store.storage.DB.Exec(updateQuery, id, req.ProfileImage, req.FullName, req.DOB, req.Gender, updatedAt)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("couldnot update the user")
	}
	return nil
}

func (store *MarchantStore) InsertProduct(userId, productId string, req *entity.AddProductReq) error {
	createdAt := time.Now()
	updatedAt := time.Now()

	ctx := context.Background()

	tx, err := store.storage.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	inventoryId, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	insertInventoryQuery := `INSERT INTO inventories
		(id, quantity, created_at, updated_at)
		VALUES ($1, $2, $3, $4);
		`
	_, err = tx.Exec(insertInventoryQuery, inventoryId.String(), req.Quantity, createdAt, updatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	insertProductQuery := `INSERT INTO products
		(id, merchent_id, category_id, product_name, description, images, size, price, inventory_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
		`
	productImages := pq.StringArray(req.ProductImages)
	productSizes := pq.Float64Array(req.Sizes)

	_, err = tx.Exec(insertProductQuery, productId, userId, req.CategoryId, req.ProductName, req.Description, productImages, productSizes, req.Price, inventoryId, createdAt, updatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (store *MarchantStore) UpdateProduct(id string, req *entity.EditProductReq) error {
	updateQuery := `
	UPDATE products
	SET product_name = $2, description = $3, images = $4, size = $5, price = $6, updated_at = $7
	WHERE id = $1;
	`
	updatedAt := time.Now()
	productImages := pq.StringArray(req.ProductImages)
	productSizes := pq.Float64Array(req.Sizes)

	res, err := store.storage.DB.Exec(updateQuery, id, req.ProductName, req.Description, productImages, productSizes, req.Price, updatedAt)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("couldnot update the products")
	}
	return nil
}

func (store *MarchantStore) DeleteProduct(productId string) error {
	deleteQuery := `
        DELETE FROM products
        USING inventories
        WHERE 
		products.id = $1
        AND 
		products.inventory_id = inventories.id;
    `

	res, err := store.storage.DB.Exec(deleteQuery, productId)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("could not delete the product")
	}
	return nil
}

func (store *MarchantStore) AddProductDiscount(req *entity.AddDiscountReq) error {
	ctx := context.Background()
	tx, err := store.storage.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	discountId, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	createdAt := time.Now()
	updatedAt := time.Now()

	if req.CardColor == "" {
		req.CardColor = "0xFF808080"
	}

	insertQuery := `
	INSERT INTO discounts
	(id, name, description, percent, card_color, start_time, end_time, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`

	res, err := tx.Exec(insertQuery, discountId.String(), req.DiscountName, req.Description, req.Percentage, req.CardColor, req.StartTime, req.EndTime, createdAt, updatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if n == 0 {
		return fmt.Errorf("couldnot insert the discounts")
	}

	updateQuery := `UPDATE products
	SET discount_id = $2, updated_at = $3 
	WHERE id = $1;
	`
	res, err = tx.Exec(updateQuery, req.ProductId, discountId.String(), updatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	n, err = res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if n == 0 {
		return fmt.Errorf("couldnot update the products")
	}
	tx.Commit()
	return nil
}

func (store *MarchantStore) EditProductDiscount(discountId string, req *entity.EditDiscountReq) error {
	updateQuery := `
	UPDATE discounts
	SET name = COALESCE($2, name),
		description = COALESCE($3, description),
		percent = COALESCE($4, percent),
		card_color = COALESCE($5, card_color),
		start_time = COALESCE($6, start_time),
		end_time = COALESCE($7, end_time),
		updated_at = $8
	WHERE id = $1;
	`
	updatedAt := time.Now()
	res, err := store.storage.DB.Exec(updateQuery, discountId, req.DiscountName, req.Description, req.Percentage, req.CardColor, req.StartTime, req.EndTime, updatedAt)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("couldnot update the discount")
	}
	return nil
}
