package dto

type UpdateProductDTO struct {
	Name              string  `json:"name"`
	Code              string  `json:"code"`
	ProductCategoryId int64   `json:"productCategoryId"`
	Description       string  `json:"description"`
	Price             float32 `json:"price"`
	Stock             int64   `json:"stock"`
	ImageUrl          string  `json:"imageurl"`
}
