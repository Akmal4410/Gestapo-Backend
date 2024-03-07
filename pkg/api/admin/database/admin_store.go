package database

import (
	"time"

	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/api/admin/database/entity"
	"github.com/google/uuid"
)

type AdminStore struct {
	storage *database.Storage
}

func NewAdminStore(storage *database.Storage) *AdminStore {
	return &AdminStore{storage: storage}
}

func (store *AdminStore) CheckCategoryExist(category string) (bool, error) {
	checkQuery := `SELECT * FROM categories WHERE category_name = $1;`
	res, err := store.storage.DB.Exec(checkQuery, category)
	if err != nil {
		return false, err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return result != 0, nil
}

func (store AdminStore) InsertCategory(req *entity.InsertCategoryReq) error {
	createdAt := time.Now()
	updatedAt := time.Now()

	insertQuery := `
	INSERT INTO categories (id, category_name, created_at, updated_at)
	VALUES ($1, $2, $3, $4);
	`

	uuId, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	_, err = store.storage.DB.Exec(insertQuery, uuId, req.Category_Name, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (store *AdminStore) GetCategories() ([]entity.GetCategoriesRes, error) {

	var categories []entity.GetCategoriesRes

	selectQuery := `
	SELECT id, category_name 
	FROM categories 
	ORDER BY category_name ASC;
	`

	rows, err := store.storage.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category entity.GetCategoriesRes
		err := rows.Scan(&category.ID, &category.Category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
func (store *AdminStore) GetUsers() ([]entity.GetUserRes, error) {
	var users []entity.GetUserRes

	selectQuery := `
	SELECT id, profile_image, full_name, user_name, phone, email, dob, gender, user_type 
	FROM user_data
	WHERE user_type != 'ADMIN'
	ORDER BY full_name ASC;
	`

	rows, err := store.storage.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.GetUserRes
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
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}
