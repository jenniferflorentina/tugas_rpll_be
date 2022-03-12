package seeds

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/transaction/domain/model"
)

func (s *Seed) SeedTransaction() {
	var transactions []model.Transaction
	var count int64
	db.Orm.Model(&transactions).Count(&count)

	if count > 0 {
		return
	}

	transactions = make([]model.Transaction, 5)
	transactions[0] = model.Transaction{
		MemberId: 1,
		Status:   2,
	}
	transactions[1] = model.Transaction{
		MemberId: 2,
		Status:   2,
	}
	transactions[2] = model.Transaction{
		MemberId: 3,
		Status:   2,
	}
	transactions[3] = model.Transaction{
		MemberId: 2,
		Status:   2,
	}
	transactions[4] = model.Transaction{
		MemberId: 4,
		Status:   2,
	}

	_ = db.Orm.Create(&transactions)
}
