package model

import (
	"HarapanBangsaMarket/base"
)

type Transaction struct {
	base.Model `gorm:"extends"`
	Id         int64
	//TODO : add FK to Member
	MemberId                   int64
	Status                     Status                     `gorm:"smallint"`
	TransactionDigitalDetailId []TransactionDigitalDetail `gorm:"ForeignKey:TransactionId;references:Id"`
}

func init() {
}
