package entity

type InsertCategoryReq struct {
	Category_Name string `json:"category_name" validate:"required"`
}
