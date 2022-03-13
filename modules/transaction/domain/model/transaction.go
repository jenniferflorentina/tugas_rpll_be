package model

import (
	"HarapanBangsaMarket/base"
	item "HarapanBangsaMarket/modules/member/domain/model"
	pay "HarapanBangsaMarket/modules/payment/domain/model"
)

type Transaction struct {
	base.Model `gorm:"extends"`
	Id         int64
	MemberId   *int64
	Status     Status              `gorm:"smallint"`
	Details    []TransactionDetail `gorm:"ForeignKey:TransactionId;references:Id"`
	Member     item.Member         `gorm:"foreignKey:MemberId"`
	Payment    pay.Payment         `gorm:"foreignKey:TransactionId;references:Id"`
}

func init() {
}
