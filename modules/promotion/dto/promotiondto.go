package dto

import (
	"HarapanBangsaMarket/base"
)

type PromotionDTO struct {
	base.DTO
	Type     string `json:"type"`
	ValidUntil string`json:"validUntil"`
}