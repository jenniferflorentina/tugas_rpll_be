package dto

type CreateUpdatePromotionDTO struct {
	Type       string                           `json:"type"`
	ValidUntil string                           `json:"validUntil"`
	Details    []CreateUpdatePromotionDetailDTO `json:"details"`
}
