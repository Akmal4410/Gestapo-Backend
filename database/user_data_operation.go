package database

import (
	"fmt"
	"time"

	"github.com/akmal4410/gestapo/models"
	"github.com/akmal4410/gestapo/services/password"
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
