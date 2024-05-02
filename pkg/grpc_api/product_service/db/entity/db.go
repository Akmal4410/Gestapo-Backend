package entity

type GetProductRes struct {
	ID            string     `json:"id"`
	MerchantID    string     `json:"merchant_id,omitempty"`
	ProductImages []string   `json:"product_images,omitempty"`
	ProductName   *string    `json:"product_name"`
	Description   *string    `json:"description,omitempty"`
	CategoryName  *string    `json:"category_name,omitempty"`
	Size          *[]float64 `json:"size,omitempty"`
	Price         float64    `json:"price,omitempty"`
	DiscountPrice *float64   `json:"discount_price,omitempty"`
}
