package internal

import "gorm.io/gorm"

type Banner struct {
	gorm.Model
	Description string `json:"description"`
	LinkTo      string `json:"link_to"`
	Priority    uint` json:"priority"`
	Attachment  Attachment `gorm:"polymorphic:Attachment;"`
}
