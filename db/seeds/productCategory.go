package seeds

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/model/product"
)

func (s *Seed) SeedProductCategory() {
	var productCategory []model.ProductCategory
	var count int64
	db.Orm.Model(&productCategory).Count(&count)

	if count > 0 {
		return
	}

	productCategory = make([]model.ProductCategory, 5)
	productCategory[0] = model.ProductCategory{
		Name:        "Digital",
	}
	productCategory[1] = model.ProductCategory{
		Name:        "Makanan",
	}
	productCategory[2] = model.ProductCategory{
		Name:        "Minuman",
	}
	productCategory[3] = model.ProductCategory{
		Name:        "Kebutuhan Rumah Tangga",
	}
	productCategory[4] = model.ProductCategory{
		Name:        "Kecantikan",
	}

	_ = db.Orm.Create(&productCategory)
}