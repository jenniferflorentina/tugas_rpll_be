package dto

type CreatePaymentDTO struct {
	Amount          float32 `json:"amount"`
	CustomerCode    string  `json:"customerCode" validate:"empty=false"`
	PaymentMethodId int64   `json:"paymentMethodId"`
	Point           float32 `json:"point"`
	TransactionId   int64   `json:"transactionId"`
}
