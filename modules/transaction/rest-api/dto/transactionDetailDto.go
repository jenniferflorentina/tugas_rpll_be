package dto

import (
	"HarapanBangsaMarket/base"
	product "HarapanBangsaMarket/modules/product/rest-api/dto"
	promotion "HarapanBangsaMarket/modules/promotion/rest-api/dto"
)

type TransactionDetailDTO struct {
	base.DTO
	Product   product.ProductDTO     `json:"product"`
	Promotion promotion.PromotionDTO `json:"promotion"`
	Quantity  int64                  `json:"quantity"`
}
