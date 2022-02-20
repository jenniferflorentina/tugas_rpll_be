package model

import (
	"HarapanBangsaMarket/base"
	item "HarapanBangsaMarket/modules/model/product"
)

type TransactionDetail struct {
	base.Model    `gorm:"extends"`
	Id            int64
	TransactionId int64
	ProductId     int64
	PromotionId		int64
	Quantity      int64
	Product       item.Product `gorm:"foreignKey:ProductId"`	
}

func init() {
}
