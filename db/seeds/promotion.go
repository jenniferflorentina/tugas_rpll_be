package seeds

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/promotion/model"
)

func (s *Seed) SeedPromotion() {
	var promotion []model.Promotion
	var count int64
	db.Orm.Model(&promotion).Count(&count)

	if count > 0 {
		return
	}

	promotions := make([]model.Promotion, 4)
	promotions[0] = model.Promotion{
		Type: "Pembelian Rp25.000",
		ValidUntil: "20 Maret 2022",

	}
	promotions[1] = model.Promotion{
		Type: "Pembelian Rp50.000",
		ValidUntil: "20 Maret 2022",

	}
	promotions[2] = model.Promotion{
		Type: "Pembelian Rp75.000",
		ValidUntil: "20 Maret 2022",

	}
	promotions[3] = model.Promotion{
		Type: "Pembelian Rp100.000",
		ValidUntil: "20 Maret 2022",

	}

	_ = db.Orm.Create(&promotions)
}