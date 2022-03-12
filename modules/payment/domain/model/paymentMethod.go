package model

import (
	"HarapanBangsaMarket/base"
)

type PaymentMethod struct {
	base.Model `gorm:"extends"`
	Id         int64
	Name       string    `gorm:"varchar(100)"`
	Code       string    `gorm:"varchar(100)"`
	Payment    []Payment `gorm:"ForeignKey:PaymentMethodId;references:Id"`
}

func init() {
}
