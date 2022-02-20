package model

import (
	"HarapanBangsaMarket/base"
	item "HarapanBangsaMarket/modules/model/product"
)

type PromotionDetail struct {
	base.Model  `gorm:"extends"`
	Id          int64
	Discount    float32 `gorm:"double"`
	PromotionId int64
	ProductId   int64
	Product     item.Product
}

func init() {
}
