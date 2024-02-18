package database

import (
	"fmt"

	"github.com/akmal4410/gestapo/services/password"
)

func (storage *Storage) CheckUserExist(column, value string) (bool, error) {
	checkQuery := fmt.Sprintf(`SELECT * FROM user_data WHERE %s = $1;`, column)
	res, err := storage.db.Exec(checkQuery, value)
	if err != nil {
		return false, err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return result != 0, nil
}

func (storage *Storage) CheckPassword(userName, pass string) (bool, error) {
	var hashPassword string
	checkQuery := `SELECT password FROM user_data WHERE user_name = $1`
	rows := storage.db.QueryRow(checkQuery, userName)
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
