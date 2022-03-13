package dto

type CreateTransactionDTO struct {
	MemberId           int64                        `json:"memberId"`
	TransactionDetails []CreateTransactionDetailDTO `json:"details"`
}
