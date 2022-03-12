package seeds

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/promotion/domain/model"
)

func (s *Seed) SeedPromotionDetail() {
	var promotionDetail []model.PromotionDetail
	var count int64
	db.Orm.Model(&promotionDetail).Count(&count)

	if count > 0 {
		return
	}

	promotionDetails := make([]model.PromotionDetail, 3)
	promotionDetails[0] = model.PromotionDetail{
		Discount:  5000,
		PromotionId: 1,
		ProductId:	11,
	}
	promotionDetails[1] = model.PromotionDetail{
		Discount:  0,
		PromotionId: 3,
		ProductId:	5,
	}
	promotionDetails[2] = model.PromotionDetail{
		Discount:  7500,
		PromotionId: 2,
		ProductId:	12,
	}
	_ = db.Orm.Create(&promotionDetails)
}