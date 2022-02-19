package db

import (
	modelUser "HarapanBangsaMarket/modules/model/user"
	modelPromotion "HarapanBangsaMarket/modules/model/promotion"
	modelProduct "HarapanBangsaMarket/modules/model/product"
)

func migrate() {
	_ = Orm.AutoMigrate(new(modelUser.User))
	_ = Orm.AutoMigrate(new(modelPromotion.Promotion))
	_ = Orm.AutoMigrate(new(modelProduct.ProductCategory))
	_ = Orm.AutoMigrate(new(modelProduct.Product))
}
