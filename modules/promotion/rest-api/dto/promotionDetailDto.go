package dto

import (
	"HarapanBangsaMarket/base"
	dto "HarapanBangsaMarket/modules/product/rest-api/dto"
)

type PromotionDetailDTO struct {
	base.DTO
	Discount  float32        `json:"discount"`
	Product   dto.ProductDTO `json:"product"`
	Promotion PromotionDTO   `json:"promotion"`
}
