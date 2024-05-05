package db

import (
	"context"
	"fmt"
	"time"

	"github.com/akmal4410/gestapo/internal/database"
	product_entity "github.com/akmal4410/gestapo/pkg/grpc_api/product_service/db/entity"
	"github.com/akmal4410/gestapo/pkg/grpc_api/user_service/db/entity"
	"github.com/akmal4410/gestapo/pkg/utils"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type UserStore struct {
	storage *database.Storage
}

func NewUserStore(storage *database.Storage) *UserStore {
	return &UserStore{
		storage: storage,
	}
}

func (store *UserStore) CheckDataExist(table, column, value string) (bool, error) {
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

func (store *UserStore) GetDiscount() (*entity.DiscountRes, error) {
	selectQuery := `
	SELECT 
    p.id AS product_id,
    d.name AS name,
	d.description AS description,
    d.percent AS percent,
    p.images[1] AS image
	FROM products p
	JOIN discounts d ON p.discount_id = d.id
	WHERE d.end_time > NOW()
	ORDER BY d.percent DESC
	LIMIT 1;
	`

	rows := store.storage.DB.QueryRow(selectQuery)
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	var discount entity.DiscountRes
	err := rows.Scan(
		&discount.ProductID,
		&discount.Name,
		&discount.Description,
		&discount.Percentage,
		&discount.ProductImage,
	)
	if err != nil {
		return nil, err
	}
	return &discount, nil
}

func (store *UserStore) GetMerchants() ([]entity.MerchantRes, error) {
	selectQuery := `
	SELECT id, full_name, profile_image 
	FROM user_data
	WHERE user_type = $1
	LIMIT 7;
	`
	rows, err := store.storage.DB.Query(selectQuery, utils.MERCHANT)
	if err != nil {
		return nil, err
	}
	var merchants []entity.MerchantRes
	defer rows.Close()
	for rows.Next() {
		var merchant entity.MerchantRes

		err := rows.Scan(
			&merchant.MerchantID,
			&merchant.Name,
			&merchant.ImageURL,
		)
		if err != nil {
			return nil, err
		}
		merchants = append(merchants, merchant)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return merchants, nil
}

func (store *UserStore) AlreadyInWishlist(req *entity.AddRemoveWishlistReq) (bool, error) {
	checkQuery := `
	SELECT user_id, product_id FROM wishlists
	WHERE user_id = $1 AND product_id = $2;
	`
	res, err := store.storage.DB.Exec(checkQuery, req.UserID, req.ProductID)
	if err != nil {
		return false, err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return result != 0, nil
}

func (store *UserStore) AddToWishlist(req *entity.AddRemoveWishlistReq) error {
	createdAt := time.Now()
	updatedAt := time.Now()

	uuId, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	insertQuery := `
	INSERT INTO wishlists (id, user_id, product_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5);
	`
	_, err = store.storage.DB.Exec(insertQuery, uuId, req.UserID, req.ProductID, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (store *UserStore) RemoveFromWishlist(req *entity.AddRemoveWishlistReq) error {
	deleteQuery := `
        DELETE FROM wishlists
        WHERE user_id = $1 AND product_id = $2;
    `

	res, err := store.storage.DB.Exec(deleteQuery, req.UserID, req.ProductID)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("could not delete the product from wishlist")
	}
	return nil
}

func (store *UserStore) GetWishlistProducts(userId string) ([]product_entity.GetProductRes, error) {
	var products []product_entity.GetProductRes

	selectQuery := `
	SELECT
    p.id AS id,
    p.product_name AS product_name,
	p.images AS product_images,
    p.price AS price
	FROM
    products p
	LEFT JOIN
    wishlists w ON p.id = w.product_id
	WHERE w.user_id = $1;
	`

	rows, err := store.storage.DB.Query(selectQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product product_entity.GetProductRes
		var images pq.StringArray

		err := rows.Scan(
			&product.ID,
			&product.ProductName,
			&images,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}
		product.ProductImages = []string(images)
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return []product_entity.GetProductRes{}, nil
	}

	return products, nil
}

func (store *UserStore) CreateUserCart(req *entity.AddToCartReq) error {
	createdAt := time.Now()
	updatedAt := time.Now()

	uuId, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	insertQuery := `
	INSERT INTO carts (id, user_id, price, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5);
	`
	_, err = store.storage.DB.Exec(insertQuery, uuId, req.UserID, 0, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (store *UserStore) AddToCard(req *entity.AddToCartReq) error {
	createdAt := time.Now()
	updatedAt := time.Now()

	ctx := context.Background()

	tx, err := store.storage.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	uuId, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	insertProductQuery := `
        INSERT INTO cart_items
        (id, cart_id, product_id, inventory_id, quantity, price, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
        `

	_, err = tx.Exec(insertProductQuery, uuId, req.CartID, req.ProductID, req.InventoryID, req.Quantity, req.Price, createdAt, updatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Calculate total price of items in the cart
	totalPriceQuery := `
        SELECT SUM(quantity * price) FROM cart_items WHERE cart_id = $1;
    `
	var totalPrice float64
	err = tx.QueryRow(totalPriceQuery, req.CartID).Scan(&totalPrice)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Update price in the carts table
	updateQuery := `
        UPDATE carts
        SET price = $1, updated_at = $2
        WHERE id = $3;
    `
	_, err = tx.Exec(updateQuery, totalPrice, updatedAt, req.CartID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (store *UserStore) GetCartID(userID string) (string, error) {
	selectQuery := `SELECT id FROM carts WHERE user_id = $1;`

	rows := store.storage.DB.QueryRow(selectQuery, userID)
	if rows.Err() != nil {
		return "", rows.Err()
	}
	var cartID string
	err := rows.Scan(&cartID)
	if err != nil {
		return "", err
	}
	return cartID, nil
}

func (store *UserStore) GetInventoryID(productID string, size float64) (string, error) {
	selectQuery := `SELECT id FROM inventories WHERE product_id = $1 AND size = $2;`

	rows := store.storage.DB.QueryRow(selectQuery, productID, size)
	if rows.Err() != nil {
		return "", rows.Err()
	}
	var inventoryID string
	err := rows.Scan(&inventoryID)
	if err != nil {
		return "", err
	}
	return inventoryID, nil
}

func (store *UserStore) GetCartItems(userId string) ([]*entity.CartItemRes, error) {
	var products []*entity.CartItemRes

	selectQuery := `
	SELECT
    p.id AS id,
	p.images AS product_images,
	p.product_name AS product_name,
	i.size AS size,
	ci.quantity AS quantity,
	ci.price AS price,
	ci.id AS cart_item_id
	FROM
    products p
	LEFT JOIN
    inventories i ON p.id = i.product_id
	LEFT JOIN
    cart_items ci ON p.id = ci.product_id
	LEFT JOIN
    carts c ON ci.cart_id = c.id
	WHERE c.user_id = $1 AND ci.inventory_id = i.id;
	`
	rows, err := store.storage.DB.Query(selectQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.CartItemRes
		var images pq.StringArray

		err := rows.Scan(
			&product.ProductID,
			&images,
			&product.Name,
			&product.Size,
			&product.Quantity,
			&product.Price,
			&product.CartItemID,
		)
		if err != nil {
			return nil, err
		}
		if len(images) > 0 {
			product.ImageURL = images[0]
		}
		products = append(products, &product)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (store *UserStore) CanDeleteCartItem(cartItemId, userId string) (bool, error) {
	query := `
        SELECT COUNT(ci.id)
        FROM cart_items ci
        JOIN carts c ON ci.cart_id = c.id
        WHERE ci.id = $1 AND c.user_id = $2
    `
	var count int
	err := store.storage.DB.QueryRow(query, cartItemId, userId).Scan(&count)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return false, err
	}
	return count > 0, nil
}

func (store *UserStore) RemoveFromCart(cartItemId, userId string) error {
	deleteQuery := `
        DELETE FROM cart_items
        WHERE id = $1
		AND cart_id IN (SELECT id FROM carts WHERE user_id = $2);
    `

	res, err := store.storage.DB.Exec(deleteQuery, cartItemId, userId)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("not deleted")
	}
	return nil
}
