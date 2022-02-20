package model

import (
	"HarapanBangsaMarket/base"
)

type Transaction struct {
	base.Model `gorm:"extends"`
	Id         int64
	MemberId   int64
	Status     Status `gorm:"smallint"`
}

func init() {
}
