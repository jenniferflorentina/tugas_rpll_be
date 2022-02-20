package model

import (
	"HarapanBangsaMarket/base"
)

type TransactionDetail struct {
	base.Model    `gorm:"extends"`
	Id            int64
	TransactionId int64
	ProductId     int64
	PromotionId		int64
	Quantity      int64
	
}

func init() {
}
