package dto

import (
	"HarapanBangsaMarket/base"
)

type ProductDTO struct {
	base.DTO
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock    int64   `json:"stock"`
	ImageUrl    string  `json:"imageurl"`
	ProductCategoryId	int64	`json:"productCategoryId"`
	ProductCategory	*ProductCategoryDTO	`json:"productCategory"`
}