package dto

type CreateTransactionDetailDTO struct {
	ProductId   int64 `json:"productId"`
	PromotionId int64 `json:"promotionId"`
	Quantity    int64 `json:"quantity"`
}
