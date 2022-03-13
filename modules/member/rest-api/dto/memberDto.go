package dto

import (
	"HarapanBangsaMarket/base"
)

type MemberDTO struct {
	base.DTO
	Name      string `json:"name" validate:"empty=false"`
	Point     int64  `json:"point" validate:"empty=false"`
	Telephone string `json:"telephone" validate:"empty=false"`
}
