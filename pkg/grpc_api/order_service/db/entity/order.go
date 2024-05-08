package entity

type CreateOrderReq struct {
	AddressID     string  `json:"address_id" validate:"required"`
	CartID        string  `json:"cart_id" validate:"required"`
	PromoID       *string `json:"promo_id"`
	Amount        float64 `json:"amount" validate:"required"`
	PaymentMode   string  `json:"action" validate:"payment_mode"`
	UserID        string  `json:"user_id"`
	TransactionID *string `json:"transaction_id"`
}
