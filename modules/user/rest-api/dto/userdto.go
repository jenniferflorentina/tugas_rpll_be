package dto

import (
	"HarapanBangsaMarket/base"
)

type UserDTO struct {
	base.DTO
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     int    `json:"role"`
}
