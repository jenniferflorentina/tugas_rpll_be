package base

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	CreatedAt time.Time `gorm:"created"`
	CreatedBy int64
	UpdatedAt time.Time `gorm:"updated"`
	UpdatedBy int64
	DeletedAt gorm.DeletedAt
	DeletedBy int64
}
