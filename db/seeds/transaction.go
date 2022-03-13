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
		MemberId: func() *int64 { i := int64(1); return &i }(),
		Status:   2,
	}
	transactions[1] = model.Transaction{
		MemberId: func() *int64 { i := int64(2); return &i }(),
		Status:   2,
	}
	transactions[2] = model.Transaction{
		MemberId: func() *int64 { i := int64(3); return &i }(),
		Status:   2,
	}
	transactions[3] = model.Transaction{
		MemberId: func() *int64 { i := int64(2); return &i }(),
		Status:   2,
	}
	transactions[4] = model.Transaction{
		MemberId: func() *int64 { i := int64(4); return &i }(),
		Status:   2,
	}

	_ = db.Orm.Create(&transactions)
}
