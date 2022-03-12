package db

import (
	modelMember "HarapanBangsaMarket/modules/member/domain/model"
	modelProduct "HarapanBangsaMarket/modules/product/domain/model"
	modelPromotion "HarapanBangsaMarket/modules/promotion/domain/model"
	modelTransaction "HarapanBangsaMarket/modules/transaction/domain/model"
	modelUser "HarapanBangsaMarket/modules/user/domain/model"
	modelPayment "HarapanBangsaMarket/modules/payment/domain/model"
)

func migrate() {
	_ = Orm.AutoMigrate(new(modelUser.User))
	_ = Orm.AutoMigrate(new(modelPromotion.Promotion))
	_ = Orm.AutoMigrate(new(modelProduct.ProductCategory))
	_ = Orm.AutoMigrate(new(modelProduct.Product))
	_ = Orm.AutoMigrate(new(modelPromotion.PromotionDetail))
	_ = Orm.AutoMigrate(new(modelTransaction.Transaction))
	_ = Orm.AutoMigrate(new(modelTransaction.TransactionDetail))
	_ = Orm.AutoMigrate(new(modelMember.Member))
	_ = Orm.AutoMigrate(new(modelPayment.PaymentMethod))
	_ = Orm.AutoMigrate(new(modelPayment.Payment))

}
