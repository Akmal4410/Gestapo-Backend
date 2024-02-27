package database

import (
	"fmt"

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

func (store *MarchantStore) GetProfile(userId string) (*entity.UserData, error) {
	selectQuery := `SELECT id, profile_image, full_name, user_name, phone, email, dob, gender, user_type FROM user_data WHERE id = $1;`
	rows := store.storage.DB.QueryRow(selectQuery, userId)
	if rows.Err() != nil {
		return &entity.UserData{}, rows.Err()
	}
	var user entity.UserData
	err := rows.Scan(
		&user.ID,
		&user.Profile_Image,
		&user.Full_Name,
		&user.User_Name,
		&user.Phone,
		&user.Email,
		&user.DOB,
		&user.Gender,
		&user.User_type,
	)
	if err != nil {
		return &entity.UserData{}, err
	}
	return &user, nil
}
