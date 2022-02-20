package dto

import (
	"HarapanBangsaMarket/base"
)

type ProductCategoryDTO struct {
	base.DTO
	Name string `json:"name"`
}