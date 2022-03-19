package service

import (
	"HarapanBangsaMarket/modules/payment/domain/model"
	"HarapanBangsaMarket/modules/payment/repository"
	"HarapanBangsaMarket/modules/payment/rest-api/dto"
)

func FindAllPaymentMethod() (*[]model.PaymentMethod, error) {
	return repository.FindAllPaymentMethod()
}

func FindOnePaymentMethod(id int64) (*model.PaymentMethod, error) {
	return repository.FindOnePaymentMethod(id)
}

func CreatePaymentMethod(paymentMethod *model.PaymentMethod) error {
	return repository.CreatePaymentMethod(paymentMethod)
}

func UpdatePaymentMethod(updateDto *dto.CreateUpdatePaymentMethodDTO, id int64) (*model.PaymentMethod, error) {
	paymentMethod, err := repository.FindOnePaymentMethod(id)
	if err != nil {
		return nil, err
	}
	if updateDto.Name != "" {
		paymentMethod.Name = updateDto.Name
	}

	if updateDto.Code != "" {
		paymentMethod.Code = updateDto.Code
	}

	paymentMethod, err = repository.UpdatePaymentMethod(paymentMethod)
	if err != nil {
		return nil, err
	}
	return paymentMethod, nil
}
