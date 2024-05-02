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
