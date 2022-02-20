package dto

type CreateProductDTO struct {
	Name        string  `json:"name" validate:"empty=false"`
	Code        string  `json:"code" validate:"empty=false"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock    int64   `json:"stock"`
	ImageUrl    string  `json:"imageurl" validate:"empty=false"`
	ProductCategoryId	int64	`json:"productCategoryId"`
}