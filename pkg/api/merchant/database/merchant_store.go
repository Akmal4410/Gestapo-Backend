package database

import (
	"context"
	"fmt"
	"time"

	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/api/merchant/database/entity"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type MarchantStore struct {
	storage *database.Storage
}

func NewMarchantStore(storage *database.Storage) *MarchantStore {
	return &MarchantStore{storage: storage}

}

func (store *MarchantStore) CheckUserExist(id, value string) (bool, error) {
	checkQuery := fmt.Sprintf(`SELECT * FROM user_data WHERE %s = $1;`, id)
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
		return &entity.GetMerchantRes{}, rows.Err()
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
		return &entity.GetMerchantRes{}, err
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

func (store *MarchantStore) InsertProduct(id string, req *entity.AddProductReq) error {
	createdAt := time.Now()
	updatedAt := time.Now()

	ctx := context.Background()

	tx, err := store.storage.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	var discountId string = ""
	if req.DiscountName != "" {
		uuId, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		discountId = uuId.String()
		insertQuery := `INSERT INTO discounts
		(id, name, percent, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5);
		`
		_, err = tx.Exec(insertQuery, discountId, req.DiscountName, req.Percent, createdAt, updatedAt)
		if err != nil {
			tx.Rollback()
			return err
		}
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
		(id, category_id, product_name, description, images, size, price, inventory_id, discount_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
		`
	productImages := pq.StringArray(req.ProductImages)
	productSizes := pq.Float64Array(req.Sizes)

	_, err = tx.Exec(insertProductQuery, id, req.CategoryId, req.ProductName, req.Description, productImages, productSizes, req.Price, inventoryId, discountId, createdAt, updatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
