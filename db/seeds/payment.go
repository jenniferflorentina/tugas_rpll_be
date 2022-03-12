package seeds

import (
	"HarapanBangsaMarket/db"
	model "HarapanBangsaMarket/modules/payment/domain/model"
)

func (s *Seed) SeedPayment() {
	var payment []model.Payment
	var count int64
	db.Orm.Model(&payment).Count(&count)

	if count > 0 {
		return
	}

	payment = make([]model.Payment, 5)
	payment[0] = model.Payment{
		Amount:          34000,
		CustomerCode:    "-",
		PaymentMethodId: 1,
		Status:          "Complete",
		Point:           0,
		TransactionId:   1,
	}
	payment[1] = model.Payment{
		Amount:          101000,
		CustomerCode:    "43729080100",
		PaymentMethodId: 2,
		Status:          "Complete",
		Point:           0,
		TransactionId:   2,
	}
	payment[2] = model.Payment{
		Amount:          73000,
		CustomerCode:    "08239849283",
		PaymentMethodId: 3,
		Status:          "Complete",
		Point:           5000,
		TransactionId:   3,
	}
	payment[3] = model.Payment{
		Amount:          50500,
		CustomerCode:    "42989409099",
		PaymentMethodId: 2,
		Status:          "Complete",
		Point:           0,
		TransactionId:   4,
	}
	payment[4] = model.Payment{
		Amount:          180000,
		CustomerCode:    "0892839289923",
		PaymentMethodId: 3,
		Status:          "Complete",
		Point:           10000,
		TransactionId:   5,
	}

	_ = db.Orm.Create(&payment)
}
