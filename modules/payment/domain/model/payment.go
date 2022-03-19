package model

import (
	"HarapanBangsaMarket/base"
)

type Payment struct {
	base.Model      `gorm:"extends"`
	Id              int64
	Amount          float32 `gorm:"double"`
	CustomerCode    string  `gorm:"varchar(100)"`
	PaymentMethodId int64
	PaymentMethod   PaymentMethod `gorm:"ForeignKey:PaymentMethodId"`
	Point           float32
	Status          string `gorm:"varchar(100)"`
	TransactionId   int64
}

func init() {
}
