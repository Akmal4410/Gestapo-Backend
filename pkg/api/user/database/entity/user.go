package entity

type GetHomeRes struct {
	Discount *DiscountRes `json:"discount,omitempty"`
}

type DiscountRes struct {
	ProductID    string  `json:"product_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Percentage   float64 `json:"percentage"`
	ProductImage string  `json:"product_image"`
	CardColor    uint32  `json:"card_color"`
}
