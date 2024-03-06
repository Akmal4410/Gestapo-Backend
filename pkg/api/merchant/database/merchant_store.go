package database

import (
	"fmt"
	"time"

	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/api/merchant/database/entity"
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

func (store *MarchantStore) GetProfile(userId string) (*entity.MerchantRes, error) {
	selectQuery := `SELECT id, profile_image, full_name, user_name, phone, email, dob, gender, user_type FROM user_data WHERE id = $1;`
	rows := store.storage.DB.QueryRow(selectQuery, userId)
	if rows.Err() != nil {
		return &entity.MerchantRes{}, rows.Err()
	}
	var user entity.MerchantRes
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
		return &entity.MerchantRes{}, err
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
