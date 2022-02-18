package db

import (
	model "HarapanBangsaMarket/modules/user/model"
	modelPromotion "HarapanBangsaMarket/modules/promotion/model"
)

func migrate() {
	_ = Orm.AutoMigrate(new(model.User))
	_ = Orm.AutoMigrate(new(modelPromotion.Promotion))
}
