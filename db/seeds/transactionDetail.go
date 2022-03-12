package seeds

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/transaction/domain/model"
)

func (s *Seed) SeedTransactionDetail() {
	var transactionDetails []model.TransactionDetail
	var count int64
	db.Orm.Model(&transactionDetails).Count(&count)

	if count > 0 {
		return
	}

	transactionDetails = make([]model.TransactionDetail, 12)
	transactionDetails[0] = model.TransactionDetail{
		TransactionId: 1,
		ProductId:     5,
		Quantity:      4,
		PromotionId:   func() *int64 { i := int64(3); return &i }(),
	}
	transactionDetails[1] = model.TransactionDetail{
		TransactionId: 1,
		ProductId:     6,
		Quantity:      3,
		PromotionId:   nil,
	}
	transactionDetails[2] = model.TransactionDetail{
		TransactionId: 1,
		ProductId:     7,
		Quantity:      4,
		PromotionId:   nil,
	}
	transactionDetails[3] = model.TransactionDetail{
		TransactionId: 2,
		ProductId:     3,
		Quantity:      100000,
		PromotionId:   nil,
	}
	transactionDetails[4] = model.TransactionDetail{
		TransactionId: 3,
		ProductId:     8,
		Quantity:      3,
		PromotionId:   func() *int64 { i := int64(4); return &i }(),
	}
	transactionDetails[5] = model.TransactionDetail{
		TransactionId: 3,
		ProductId:     11,
		Quantity:      2,
		PromotionId:   nil,
	}
	transactionDetails[6] = model.TransactionDetail{
		TransactionId: 3,
		ProductId:     12,
		Quantity:      2,
		PromotionId:   nil,
	}
	transactionDetails[7] = model.TransactionDetail{
		TransactionId: 4,
		ProductId:     1,
		Quantity:      50000,
		PromotionId:   nil,
	}
	transactionDetails[8] = model.TransactionDetail{
		TransactionId: 5,
		ProductId:     7,
		Quantity:      5,
		PromotionId:   nil,
	}
	transactionDetails[9] = model.TransactionDetail{
		TransactionId: 5,
		ProductId:     9,
		Quantity:      5,
		PromotionId:   nil,
	}
	transactionDetails[10] = model.TransactionDetail{
		TransactionId: 5,
		ProductId:     10,
		Quantity:      2,
		PromotionId:   nil,
	}
	transactionDetails[11] = model.TransactionDetail{
		TransactionId: 5,
		ProductId:     12,
		Quantity:      2,
		PromotionId:   nil,
	}

	_ = db.Orm.Create(&transactionDetails)
}
