package seeds

import (
	"HarapanBangsaMarket/db"
	model "HarapanBangsaMarket/modules/model/payment"
)

func (s *Seed) SeedPayment() {
	var payment []model.Payment
	var count int64
	db.Orm.Model(&payment).Count(&count)

	if count > 0 {
		return
	}

	payment = make([]model.Payment, 3)
	payment[0] = model.Payment{
		Amount:          100000,
		CustomerCode:    "Amalia",
		PaymentMethodId: 1,
		Status:          "Complete",
		TransactionId:   1,
	}
	payment[1] = model.Payment{
		Amount:          200000,
		CustomerCode:    "Bona",
		PaymentMethodId: 2,
		Status:          "Complete",
		TransactionId:   2,
	}
	payment[2] = model.Payment{
		Amount:          300000,
		CustomerCode:    "Celia",
		PaymentMethodId: 3,
		Status:          "Complete",
		TransactionId:   3,
	}

	_ = db.Orm.Create(&payment)
}
