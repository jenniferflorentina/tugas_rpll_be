package model

import (
	"HarapanBangsaMarket/base"
)

type ProductCategory struct {
	base.Model `gorm:"extends"`
	Id         int64
	Name       string    `gorm:"varchar(100)"`
	Products   []Product `gorm:"ForeignKey:ProductCategoryId;references:Id"`
}

func init() {
}
