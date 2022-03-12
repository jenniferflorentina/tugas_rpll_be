package service

import (
	dto "HarapanBangsaMarket/modules/dto/payment"
	model "HarapanBangsaMarket/modules/model/payment"
	repository "HarapanBangsaMarket/modules/repository/payment"
)

func FindPayment() (*[]model.Payment, error) {
	return repository.FindPayment()
}

func FindOnePayment(id int64) (*model.Payment, error) {
	return repository.FindOnePayment(id)
}

func CreatePayment(payments *model.Payment) error {
	return repository.CreatePayment(payments)
}

func UpdatePayment(updateDto *dto.UpdatePaymentDTO, id int64) (*model.Payment, error) {
	payments, err := repository.FindOnePayment(id)
	if err != nil {
		return nil, err
	}
	if updateDto.CustomerCode != "" {
		payments.CustomerCode = updateDto.CustomerCode
	}
	if updateDto.Status != "" {
		payments.Status = updateDto.Status
	}
	if updateDto.Amount != 0 {
		payments.Amount = updateDto.Amount
	}
	payments, err = repository.UpdatePayment(payments)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func DeletePayment(id int64) (*model.Payment, error) {
	payments, err := repository.FindOnePayment(id)
	if err != nil {
		return nil, err
	}
	return repository.DeletePayment(payments)
}
