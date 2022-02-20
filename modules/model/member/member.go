package model

import (
	"HarapanBangsaMarket/base"
)

type Member struct {
	base.Model `gorm:"extends"`
	Id         int64
	MemberID   int64
	Point      int64
	Telephone  string `gorm:"varchar(100)"`
}

func init() {
}
