package dto

type UpdatePaymentDTO struct {
	Amount        float32 `json:"amount"`
	CustomerCode  string  `json:"customerCode"`
	PaymentMethod int64   `json:"paymentMethod"`
	Status        string  `json:"status"`
	TransactionID int64   `json:"transactionID"`
}
