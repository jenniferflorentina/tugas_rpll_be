package db

import (
	model "HarapanBangsaMarket/modules/user/model"
	modelPromotion "HarapanBangsaMarket/modules/promotion/model"
	modelProduct "HarapanBangsaMarket/modules/product/model"
)

func migrate() {
	_ = Orm.AutoMigrate(new(model.User))
	_ = Orm.AutoMigrate(new(modelPromotion.Promotion))
	_ = Orm.AutoMigrate(new(modelProduct.ProductCategory))
	_ = Orm.AutoMigrate(new(modelProduct.Product))
}
