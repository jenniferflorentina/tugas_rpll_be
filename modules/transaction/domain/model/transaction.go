package model

import (
	"HarapanBangsaMarket/base"
	item "HarapanBangsaMarket/modules/member/domain/model"
)

type Transaction struct {
	base.Model `gorm:"extends"`
	Id         int64
	//TODO : add FK to Member
	MemberId                  int64
	Status                    Status                     `gorm:"smallint"`
	TransactionDetails        []TransactionDetail        `gorm:"ForeignKey:TransactionId;references:Id"`
	Member                    item.Member                `gorm:"foreignKey:MemberId"`
}

func init() {
}
