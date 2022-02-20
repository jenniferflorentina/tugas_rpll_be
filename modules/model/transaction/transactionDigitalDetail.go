package model

import (
	"HarapanBangsaMarket/base"
	item "HarapanBangsaMarket/modules/model/product"
)

type TransactionDigitalDetail struct {
	base.Model    `gorm:"extends"`
	Id            int64
	CustomerCode  string `gorm:"varchar(100)"`
	ProductId     int64
	TransactionId int64
	Transaction   Transaction
	Product       item.Product `gorm:"foreignKey:ProductId"`
}

func init() {
}
