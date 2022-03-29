package model

import (
	"HarapanBangsaMarket/base"
)

type Product struct {
	base.Model        `gorm:"extends"`
	Id                int64
	ProductCategoryId int64
	ProductCategory   ProductCategory
	Name              string  `gorm:"varchar(100)"`
	Code              string  `gorm:"varchar(100)"`
	Description       string  `gorm:"varchar(255)"`
	Price             float32 `gorm:"double"`
	Stock             int64
	ImageUrl          []byte `gorm:"column:image_url" schema:"-"`
}

func init() {
}
