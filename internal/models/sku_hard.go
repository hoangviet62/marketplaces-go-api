package internal

import (
	"time"

	"gorm.io/gorm"
)

type SkuHard struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	SpecHard  SpecHard       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
