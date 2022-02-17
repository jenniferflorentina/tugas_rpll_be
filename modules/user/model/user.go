package model

import (
	"HarapanBangsaMarket/base"
)

type User struct {
	base.Model `gorm:"extends"`
	Id         int64
	Username   string `gorm:"varchar(100)"`
	Password   string `gorm:"varchar(255)"`
	Name       string `gorm:"varchar(40)"`
	Role       Role   `gorm:"smallint"`
}

func init() {
}
