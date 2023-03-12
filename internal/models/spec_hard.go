package internal

import (
	"time"

	"gorm.io/gorm"
)

type SpecHard struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	SkuHardId uint
	SkuHard   SkuHard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
