package db

import (
	model "HarapanBangsaMarket/modules/user/model"
)

func migrate() {
	_ = Orm.AutoMigrate(new(model.User))
}
