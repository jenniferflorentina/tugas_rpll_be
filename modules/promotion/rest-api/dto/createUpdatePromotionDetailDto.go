package dto

type CreateUpdatePromotionDetailDTO struct {
	Id        int64   `json:"id"`
	Discount  float32 `json:"discount"`
	ProductId int64   `json:"productId"`
}
