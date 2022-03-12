package repository

import (
	"HarapanBangsaMarket/db"
	model "HarapanBangsaMarket/modules/payment/domain/model"
	"errors"
)

func FindPayment() (*[]model.Payment, error) {
	var payments []model.Payment
	result := db.Orm.Find(&payments)
	if result.Error != nil {
		return nil, result.Error
	}

	return &payments, nil
}

func FindOnePayment(id int64) (*model.Payment, error) {
	var payment model.Payment
	result := db.Orm.First(&payment, id)
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

func UpdatePayment(payment *model.Payment) (*model.Payment, error) {
	result := db.Orm.Save(&payment)
	if result.Error != nil {
		return nil, result.Error
	}
	payment, _ = FindOnePayment(payment.Id)
	return payment, nil
}

func DeletePayment(payment *model.Payment) (*model.Payment, error) {
	result := db.Orm.Delete(&payment)
	if result.Error != nil {
		return nil, result.Error
	}
	return payment, nil
}
