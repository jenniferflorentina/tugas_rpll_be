package seeds

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/product/domain/model"
)

func (s *Seed) SeedProduct() {
	var product []model.Product
	var count int64
	db.Orm.Model(&product).Count(&count)

	if count > 0 {
		return
	}

	product = make([]model.Product, 12)
	product[0] = model.Product{
		ProductCategoryId: 1,
		Name:              "Flazz BCA",
		Code:              "Dp-01",
		Description:       "TopUp Flazz BCA",
		Price:             500,
		Stock:             0,
	}
	product[1] = model.Product{
		ProductCategoryId: 1,
		Name:              "E-Toll",
		Code:              "Dp-02",
		Description:       "TopUp E-Toll",
		Price:             1000,
		Stock:             0,
	}
	product[2] = model.Product{
		ProductCategoryId: 1,
		Name:              "OVO",
		Code:              "Dp-03",
		Description:       "TopUp OVO",
		Price:             1000,
		Stock:             0,
	}
	product[3] = model.Product{
		ProductCategoryId: 1,
		Name:              "GoPay",
		Code:              "Dp-04",
		Description:       "TopUp GoPay",
		Price:             1000,
		Stock:             0,
	}
	product[4] = model.Product{
		ProductCategoryId: 2,
		Name:              "Indomie Goreng",
		Code:              "Pm-01",
		Description:       "makanan bermicin 85g",
		Price:             2500,
		Stock:             100,
	}
	product[5] = model.Product{
		ProductCategoryId: 2,
		Name:              "Indomie Soto Makkasar",
		Code:              "Pm-02",
		Description:       "makanan bermicin 75g",
		Price:             3000,
		Stock:             120,
	}
	product[6] = model.Product{
		ProductCategoryId: 3,
		Name:              "Coca cola",
		Code:              "Pd-01",
		Description:       "minuman soda 350ml",
		Price:             5000,
		Stock:             180,
	}
	product[7] = model.Product{
		ProductCategoryId: 3,
		Name:              "Teh Pucuk",
		Code:              "Pm-01",
		Description:       "minuman teh 450ml",
		Price:             3000,
		Stock:             99,
	}
	product[8] = model.Product{
		ProductCategoryId: 4,
		Name:              "Tessa",
		Code:              "Pr-01",
		Description:       "Tissu 500 lembar",
		Price:             18000,
		Stock:             150,
	}
	product[9] = model.Product{
		ProductCategoryId: 4,
		Name:              "S.O.S",
		Code:              "Pr-02",
		Description:       "Pembersih lantai aroma lemon",
		Price:             16500,
		Stock:             102,
	}
	product[10] = model.Product{
		ProductCategoryId: 2,
		Name:              "Ponds",
		Code:              "Pk-01",
		Description:       "Sabun pencuci wajah",
		Price:             15000,
		Stock:             101,
	}
	product[11] = model.Product{
		ProductCategoryId: 5,
		Name:              "Biorre",
		Code:              "Pk-02",
		Description:       "Sabun Mandi",
		Price:             21000,
		Stock:             100,
	}

	_ = db.Orm.Create(&product)
}
