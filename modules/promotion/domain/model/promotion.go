package model

import (
	"HarapanBangsaMarket/base"
)

type Promotion struct {
	base.Model `gorm:"extends"`
	Id         int64
	Type       string `gorm:"varchar(100)"`
	ValidUntil string `gorm:"varchar(100)"`
}

func init() {
}
