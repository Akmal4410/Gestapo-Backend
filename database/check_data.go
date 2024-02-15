package database

import "fmt"

func (storage *Storage) CheckUserExist(column, value string) (bool, error) {
	checkStmt := fmt.Sprintf(`SELECT * FROM user_data WHERE %s = $1;`, column)
	res, err := storage.db.Exec(checkStmt, value)
	if err != nil {
		return false, err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return result != 0, nil
}
