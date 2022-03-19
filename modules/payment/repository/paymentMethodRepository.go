package repository

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/payment/domain/model"
	"errors"
)

func FindAllPaymentMethod() (*[]model.PaymentMethod, error) {
	var paymentMethods []model.PaymentMethod
	result := db.Orm.Where("deleted_at is null").Find(&paymentMethods)
	if result.Error != nil {
		return nil, result.Error
	}

	return &paymentMethods, nil
}

func FindOnePaymentMethod(id int64) (*model.PaymentMethod, error) {
	var paymentMethod model.PaymentMethod
	result := db.Orm.First(&paymentMethod, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no payment method found")
	}

	return &paymentMethod, nil
}

func CreatePaymentMethod(paymentMethod *model.PaymentMethod) error {
	result := db.Orm.Create(paymentMethod)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdatePaymentMethod(paymentMethod *model.PaymentMethod) (*model.PaymentMethod, error) {
	result := db.Orm.Save(&paymentMethod)
	if result.Error != nil {
		return nil, result.Error
	}
	paymentMethod, _ = FindOnePaymentMethod(paymentMethod.Id)
	return paymentMethod, nil
}
