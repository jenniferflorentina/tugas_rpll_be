package model

import (
	"HarapanBangsaMarket/base"
	"time"
)

type Transaction struct {
	base.Model         `gorm:"extends"`
	Id                 int64
	Member_id          int64
	Created_at		   time.Date
	Created_by	       int64
	Status             Status `gorm:"smallint"`
}

func init() {
}
