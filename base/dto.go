package base

import "time"

type DTO struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int64     `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy int64     `json:"updatedBy"`
	DeletedAt time.Time `json:"deletedAt"`
	DeletedBy int64     `json:"deletedBy"`
}
