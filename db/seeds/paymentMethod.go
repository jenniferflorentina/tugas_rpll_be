package seeds

import (
	"HarapanBangsaMarket/db"
	model "HarapanBangsaMarket/modules/model/payment"
)

func (s *Seed) SeedPaymentMethod() {
	var paymentMethod []model.PaymentMethod
	var count int64
	db.Orm.Model(&paymentMethod).Count(&count)

	if count > 0 {
		return
	}

	paymentMethod = make([]model.PaymentMethod, 3)
	paymentMethod[0] = model.PaymentMethod{
		Code: "PMCash",
		Name: "Cash",
	}
	paymentMethod[1] = model.PaymentMethod{
		Code: "PMCard",
		Name: "Debit",
	}
	paymentMethod[2] = model.PaymentMethod{
		Code: "PMEmoney",
		Name: "E-Money",
	}

	_ = db.Orm.Create(&paymentMethod)
}
