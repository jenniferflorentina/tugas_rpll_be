package dto

type CreatePaymentDTO struct {
	Amount        float32 `json:"amount" validate:"empty=false"`
	CustomerCode  string  `json:"customerCode" validate:"empty=false"`
	PaymentMethod int64   `json:"paymentMethod" validate:"empty=false"`
	Status        string  `json:"status" validate:"empty=false"`
	TransactionID int64   `json:"transactionID" validate:"empty=false"`
}
