package entity

import "github.com/akmal4410/gestapo/pkg/grpc_api/product_service/db/entity"

type GetHomeRes struct {
	Discount  *DiscountRes           `json:"discount,omitempty"`
	Merchants []MerchantRes          `json:"merchants,omitempty"`
	Products  []entity.GetProductRes `json:"products,omitempty"`
}

type DiscountRes struct {
	ProductID    string  `json:"product_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Percentage   float64 `json:"percentage"`
	ProductImage string  `json:"product_image"`
	CardColor    uint32  `json:"card_color"`
}

type MerchantRes struct {
	MerchantID string  `json:"merchant_id"`
	Name       string  `json:"name"`
	ImageURL   *string `json:"image_url,omitempty"`
}

type AddRemoveWishlistReq struct {
	Action    string `json:"action" validate:"wishlist_action"`
	ProductID string `json:"product_id" validate:"required"`
	UserID    string `json:"user_id" validate:"required"`
}

type AddToCartReq struct {
	ProductID   string  `json:"product_id" validate:"required"`
	Size        float64 `json:"size" validate:"required"`
	Quantity    int32   `json:"quantity" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	UserID      string  `json:"user_id" validate:"required"`
	CartID      string  `json:"cart_id"`
	InventoryID string  `json:"inventory_id"`
}

type CartItemRes struct {
	ProductID  string  `json:"product_id"`
	CartItemID string  `json:"cart_item_id"`
	ImageURL   string  `json:"image_url"`
	Name       string  `json:"name"`
	Size       float64 `json:"size"`
	Quantity   int32   `json:"quantity"`
	Price      float64 `json:"price"`
}
