package repository

import (
	"HarapanBangsaMarket/db"
	model "HarapanBangsaMarket/modules/payment/domain/model"
	"errors"
)

func FindPayment() (*[]model.Payment, error) {
	var payments []model.Payment
	result := db.Orm.Where("deleted_at is null").Preload("PaymentMethod").Find(&payments)
	if result.Error != nil {
		return nil, result.Error
	}

	return &payments, nil
}

func FindOnePayment(id int64) (*model.Payment, error) {
	var payment model.Payment
	result := db.Orm.Preload("PaymentMethod").First(&payment, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no payment found")
	}

	return &payment, nil
}

func FindOnePaymentByTransactionId(transactionId int64) (*model.Payment, error) {
	var payment model.Payment
	result := db.Orm.Where(&model.Payment{TransactionId: transactionId}).Preload("PaymentMethod").First(&payment)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no payment found")
	}

	return &payment, nil
}

func CreatePayment(payment *model.Payment) error {
	result := db.Orm.Create(payment)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
