package model

import (
	"HarapanBangsaMarket/base"
	transaction "HarapanBangsaMarket/modules/transaction/domain/model"
	customer "HarapanBangsaMarket/modules/user/domain/model"
)

type Payment struct {
	base.Model      `gorm:"extends"`
	Id              int64
	Amount          float32 `gorm:"double"`
	CustomerCode    string  `gorm:"varchar(100)"`
	userId          customer.User
	PaymentMethodId int64
	PaymentMethod   PaymentMethod
	Status          string `gorm:"varchar(100)"`
	TransactionId   int64
	Transaction     transaction.Transaction
}

func init() {
}
