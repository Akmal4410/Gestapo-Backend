package entity

type GetHomeRes struct {
	Discount *DiscountRes `json:"discount,omitempty"`
}

type DiscountRes struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Percentage   float64 `json:"percentage"`
	ProductID    string  `json:"product_id"`
	ProductImage string  `json:"product_image"`
	CardColor    uint32  `json:"card_color"`
}
