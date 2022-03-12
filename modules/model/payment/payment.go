package model

import (
	"HarapanBangsaMarket/base"
	transaction "HarapanBangsaMarket/modules/model/transaction"
	customer "HarapanBangsaMarket/modules/model/user"
)

type Payment struct {
	base.Model      `gorm:"extends"`
	Id              int64
	Amount          float32       `gorm:"double"`
	CustomerCode    string        `gorm:"varchar(100)"`
	UserId          customer.User `gorm:"varchar(100)"`
	PaymentMethodId int64
	PaymentMethod   PaymentMethod
	Status          string `gorm:"varchar(100)"`
	TransactionId   int64
	Transaction     transaction.Transaction
}

func init() {
}
