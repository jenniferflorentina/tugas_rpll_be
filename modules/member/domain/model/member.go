package model

import (
	"HarapanBangsaMarket/base"
)

type Member struct {
	base.Model `gorm:"extends"`
	Id         int64
	Name       string `gorm:"varchar(100)"`
	Point      int64
	Telephone  string `gorm:"varchar(100);uniqueIndex;"`
}

func init() {
}
