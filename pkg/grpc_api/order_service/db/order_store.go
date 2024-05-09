package db

import (
	"context"
	"fmt"
	"time"

	"github.com/akmal4410/gestapo/internal/database"
	"github.com/akmal4410/gestapo/pkg/grpc_api/order_service/db/entity"
	user_entity "github.com/akmal4410/gestapo/pkg/grpc_api/user_service/db/entity"

	"github.com/akmal4410/gestapo/pkg/utils"
	"github.com/google/uuid"
)

type OrderStore struct {
	storage *database.Storage
}

func NewOrderStore(storage *database.Storage) *OrderStore {
	return &OrderStore{storage: storage}
}

// returns true if the user has order more than two time
func (store *OrderStore) CheckCODIsAvailable(UserID string) (bool, error) {
	selectQuery := `SELECT COUNT(user_id) FROM order_details WHERE user_id = $1;`
	var count int
	err := store.storage.DB.QueryRow(selectQuery, UserID).Scan(&count)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return false, err
	}
	return count > 2, nil
}

func (store *OrderStore) CreateOrder(req *entity.CreateOrderReq) error {
	createdAt := time.Now()
	updatedAt := time.Now()

	ctx := context.Background()
	tx, err := store.storage.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	paymentID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	var status string
	if req.PaymentMode == utils.COD {
		status = utils.PaymentPending
	} else {
		status = utils.PaymentCompleted
	}

	insertPaymentQuery := `
	INSERT INTO payment_details
	(id, amount, provider, status, transaction_id,  created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	_, err = tx.Exec(insertPaymentQuery, paymentID, req.Amount, req.PaymentMode, status, req.TransactionID, createdAt, updatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	orderDetailID, err := uuid.NewRandom()
	if err != nil {
		tx.Rollback()
		return err
	}

	insertOrderDetailQuery := `
	INSERT INTO order_details
	(id, user_id, payment_id, address_id, promo_id, amount, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`
	_, err = tx.Exec(insertOrderDetailQuery, orderDetailID, req.UserID, paymentID, req.AddressID, req.PromoID, req.Amount, createdAt, updatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	//select all the cart items
	selectOrderItemsQuery := `
	SELECT product_id, inventory_id, quantity, price 
	FROM cart_items 
	WHERE cart_id = $1;
	`
	rows, err := store.storage.DB.Query(selectOrderItemsQuery, req.CartID)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer rows.Close()

	var cartItems []*user_entity.CartItemRes
	for rows.Next() {
		var item user_entity.CartItemRes

		err := rows.Scan(
			&item.ProductID,
			&item.InventoryID,
			&item.Quantity,
			&item.Price,
		)

		if err != nil {
			tx.Rollback()
			return err
		}
		cartItems = append(cartItems, &item)
	}

	err = rows.Err()
	if err != nil {
		tx.Rollback()
		return err
	}

	var discountedPercent *float64
	if req.PromoID != nil {
		selectQuery := `SELECT percent FROM promo_codes WHERE id = $1;`
		err = tx.QueryRow(selectQuery, req.PromoID).Scan(&discountedPercent)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, item := range cartItems {
		amount := item.Price
		if discountedPercent != nil {
			amount = amount * (1 - *discountedPercent/100)
		}

		var size float32
		selectSizeQuery := `SELECT size FROM inventories WHERE id = $1;`
		err = tx.QueryRow(selectSizeQuery, item.InventoryID).Scan(&size)
		if err != nil {
			tx.Rollback()
			return err
		}

		orderItemID, err := uuid.NewRandom()
		if err != nil {
			tx.Rollback()
			return err
		}

		insertOrderItemQuery := `
		INSERT INTO order_items
		(id, order_id, product_id, size, quantity, amount, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
		`

		_, err = tx.Exec(insertOrderItemQuery, orderItemID, orderDetailID, item.ProductID, size, item.Quantity, amount, utils.OrderActive, createdAt, updatedAt)
		if err != nil {
			tx.Rollback()
			return err
		}

		//Inserting into tracking_details table
		trackingID, err := uuid.NewRandom()
		if err != nil {
			tx.Rollback()
			return err
		}

		insertTrackingQuery := `
		INSERT INTO tracking_details
		(id, order_item_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5);
		`

		_, err = tx.Exec(insertTrackingQuery, trackingID, orderItemID, utils.TrackingStatus1, createdAt, updatedAt)
		if err != nil {
			tx.Rollback()
			return err
		}

		//Inserting into tracking_items table
		trackingItemID, err := uuid.NewRandom()
		if err != nil {
			tx.Rollback()
			return err
		}

		insertTrackingItmeQuery := `
		INSERT INTO tracking_items
		(id, tracking_id, title, summary, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6);
		`

		_, err = tx.Exec(insertTrackingItmeQuery, trackingItemID, trackingID, utils.TrackingTitles[utils.TrackingStatus1], utils.TrackingSummeries[utils.TrackingStatus1], createdAt, updatedAt)
		if err != nil {
			tx.Rollback()
			return err
		}

		// Update quantity in inventories in table
		updateQuery := `
        UPDATE inventories
        SET quantity = quantity - $1, updated_at = $2
        WHERE id = $3;
    	`
		res, err := tx.Exec(updateQuery, item.Quantity, updatedAt, item.InventoryID)
		if err != nil {
			tx.Rollback()
			return err
		}
		n, err := res.RowsAffected()
		if err != nil {
			tx.Rollback()
			return err
		}
		if n == 0 {
			tx.Rollback()
			return fmt.Errorf("could update inventories")
		}
	}

	//Deleting the cart_items
	deleteCartItemsQuery := `DELETE FROM cart_items WHERE cart_id = $1;`

	res, err := store.storage.DB.Exec(deleteCartItemsQuery, req.CartID)
	if err != nil {
		tx.Rollback()
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if n == 0 {
		tx.Rollback()
		return fmt.Errorf("could not clear the cart items")
	}

	//Deleting the cart_items
	deleteCartQuery := `DELETE FROM carts WHERE id = $1;`

	res, err = store.storage.DB.Exec(deleteCartQuery, req.CartID)
	if err != nil {
		tx.Rollback()
		return err
	}
	n, err = res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if n == 0 {
		tx.Rollback()
		return fmt.Errorf("could not clear the cart")
	}

	tx.Commit()
	return nil
}
