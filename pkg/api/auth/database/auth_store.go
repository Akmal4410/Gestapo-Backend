package database

import (
	"fmt"
	"time"

	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/api/auth/database/entity"
	"github.com/akmal4410/gestapo/pkg/service/password"
	"github.com/google/uuid"
)

type AuthStore struct {
	storage *database.Storage
}

func NewAuthStore(storage *database.Storage) *AuthStore {
	return &AuthStore{storage: storage}

}

func (store AuthStore) InsertUser(user *entity.SignupReq) (id string, err error) {
	var column string
	var value string
	if user.Email != "" {
		column = "email"
		value = user.Email
	} else if user.Phone != "" {
		column = "phone"
		value = user.Phone
	}
	createdAt := time.Now()
	updatedAt := time.Now()

	user.Password, err = password.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	insertQuery := fmt.Sprintf(`
	INSERT INTO user_data (id, user_name, %s, user_type, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7);
	`, column)

	uuId, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	_, err = store.storage.DB.Exec(insertQuery, uuId, user.UserName, value, user.UserType, user.Password, createdAt, updatedAt)
	if err != nil {
		return "", err
	}
	return uuId.String(), nil
}

func (store AuthStore) ChangePassword(req *entity.ForgotPassword) (err error) {
	var column string
	var value string
	if req.Email != "" {
		column = "email"
		value = req.Email
	} else if req.Phone != "" {
		column = "phone"
		value = req.Phone
	}
	updatedAt := time.Now()

	req.Password, err = password.HashPassword(req.Password)
	if err != nil {
		return err
	}

	updateQuery := fmt.Sprintf(`UPDATE user_data SET password = $1, updated_at = $2 WHERE %s = $3`, column)
	_, err = store.storage.DB.Exec(updateQuery, req.Password, updatedAt, value)
	if err != nil {
		return err
	}
	return nil
}

type TokenPayload struct {
	UserId   string
	UserName string
	UserType string
}

func (store AuthStore) GetTokenPayload(column, value string) (*TokenPayload, error) {
	selectQuery := fmt.Sprintf(`SELECT id, user_name, user_type FROM user_data WHERE %s = $1;`, column)
	rows := store.storage.DB.QueryRow(selectQuery, value)
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	var tokenPayload TokenPayload
	err := rows.Scan(&tokenPayload.UserId, &tokenPayload.UserName, &tokenPayload.UserType)
	if err != nil {
		return nil, err
	}

	return &tokenPayload, nil
}

func (store AuthStore) CheckDataExist(column, value string) (bool, error) {
	checkQuery := fmt.Sprintf(`SELECT * FROM user_data WHERE %s = $1;`, column)
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

func (store AuthStore) CheckPassword(userName, pass string) (bool, error) {
	var hashPassword string
	checkQuery := `SELECT password FROM user_data WHERE user_name = $1`
	rows := store.storage.DB.QueryRow(checkQuery, userName)
	if rows.Err() != nil {
		return false, rows.Err()
	}

	err := rows.Scan(&hashPassword)
	if err != nil {
		return false, err
	}

	res := password.VerifyPassword(hashPassword, pass)
	return res, nil
}
