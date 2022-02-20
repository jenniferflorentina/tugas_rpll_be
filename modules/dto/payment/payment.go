package dto

import (
	"HarapanBangsaMarket/base"
)

type PaymentDTO struct {
	base.DTO
	Amount          float32           `json:"amount"`
	CustomerCode    string            `json:"customerCode"`
	PaymentMethodId int64             `json:"paymentMethodId"`
	Status          string            `json:"status"`
	TransactionId   int64             `json:"transactionID"`
	PaymentMethod   *PaymentMethodDTO `json:"paymentMethod"`
}
