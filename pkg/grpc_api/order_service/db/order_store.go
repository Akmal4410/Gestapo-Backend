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
	selectQuery := `
	SELECT COUNT(user_id) FROM order_details WHERE user_id = $1;
	`
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

	paymentID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	tx, err := store.storage.DB.BeginTx(ctx, nil)
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
	(id, amount, provider, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6);
	`

	_, err = tx.Exec(insertPaymentQuery, paymentID, req.Amount, req.PaymentMode, status, createdAt, updatedAt)
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

	//select all the order items
	selectOrderItemsQuery := `
	SELECT
    product_id, quantity, price 
	FROM order_items
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
		var item *user_entity.CartItemRes

		err := rows.Scan(
			&item.ProductID,
			&item.Quantity,
			&item.Price,
		)

		if err != nil {
			tx.Rollback()
			return err
		}
		cartItems = append(cartItems, item)
	}

	err = rows.Err()
	if err != nil {
		tx.Rollback()
		return err
	}
	//////////////////////////////
	var discountedPrice float64
	if req.PromoID != nil {
		selectQuery := `
		SELECT c.price * (1 - p.percent / 100) AS discounted_price
		FROM carts c 
		INNER JOIN promo_codes p ON p.id = $2
		WHERE c.id = $1;
		`

		err = tx.QueryRow(selectQuery, req.CartID, req.PromoID).Scan(&discountedPrice)
		if err != nil {
			fmt.Println("Error executing query:", err)
			tx.Rollback()
			return err
		}
	}

	//////////////////////////////

	for _, item := range cartItems {
		//Inserting into order_items table
		amount := item.Price
		if discountedPrice != 0 {
			amount = item.Price - (discountedPrice / float64(len(cartItems)))
		}
		orderItemID, err := uuid.NewRandom()
		if err != nil {
			tx.Rollback()
			return err
		}

		insertOrderItemQuery := `
		INSERT INTO order_items
		(id, order_id, product_id, quantity, amount, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
		`

		_, err = tx.Exec(insertOrderItemQuery, orderItemID, item.ProductID, item.Quantity, amount, utils.OrderActive, createdAt, updatedAt)
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
	}

	//Deleting the cart_items
	deleteCartItemsQuery := `
	DELETE FROM cart_items
	WHERE cart_id = $1;
	`

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
	deleteCartQuery := `
		DELETE FROM carts
		WHERE id = $1;
		`

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
