package database

import (
	"fmt"
	"time"

	"github.com/akmal4410/gestapo/internal/models"
	"github.com/akmal4410/gestapo/pkg/service/password"
)

func (storage *Storage) InsertUser(user *models.SignupReq) (err error) {
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
		return err
	}

	insertQuery := fmt.Sprintf(`
	INSERT INTO user_data (user_name, %s, user_type, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6);
	`, column)

	_, err = storage.db.Exec(insertQuery, user.UserName, value, user.UserType, user.Password, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (storage *Storage) ChangePassword(req *models.ForgotPassword) (err error) {
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
	_, err = storage.db.Exec(updateQuery, req.Password, updatedAt, value)
	if err != nil {
		return err
	}
	return nil
}

type TokenPayload struct {
	UserName string
	UserType string
}

func (storage *Storage) GetTokenPayload(column, value string) (*TokenPayload, error) {
	selectQuery := fmt.Sprintf(`SELECT user_name, userType FROM user_data WHERE %s = $1;`, column)
	rows := storage.db.QueryRow(selectQuery, value)
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	var tokenPayload TokenPayload
	err := rows.Scan(&tokenPayload)
	if err != nil {
		return nil, err
	}

	return &tokenPayload, nil
}
