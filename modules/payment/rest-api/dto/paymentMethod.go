package dto

import (
	"HarapanBangsaMarket/base"
)

type PaymentMethodDTO struct {
	base.DTO
	Code string `json:"code"`
	Name string `json:"name"`
}
