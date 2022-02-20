package db

import (
	modelProduct "HarapanBangsaMarket/modules/model/product"
	modelPromotion "HarapanBangsaMarket/modules/model/promotion"
	modelTransaction "HarapanBangsaMarket/modules/model/transaction"
	modelUser "HarapanBangsaMarket/modules/model/user"
)

func migrate() {
	_ = Orm.AutoMigrate(new(modelUser.User))
	_ = Orm.AutoMigrate(new(modelPromotion.Promotion))
	_ = Orm.AutoMigrate(new(modelProduct.ProductCategory))
	_ = Orm.AutoMigrate(new(modelProduct.Product))
	_ = Orm.AutoMigrate(new(modelPromotion.PromotionDetail))
	_ = Orm.AutoMigrate(new(modelTransaction.Transaction))
	_ = Orm.AutoMigrate(new(modelTransaction.TransactionDigitalDetail))
}
