package internal

import "gorm.io/gorm"

type Banner struct {
	gorm.Model
	Description string
	LinkTo      string
	Priority    uint
	Attachment  Attachment `gorm:"polymorphic:Attachment;"`
}
