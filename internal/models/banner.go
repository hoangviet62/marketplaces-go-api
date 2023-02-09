package internal

import "gorm.io/gorm"

type Banner struct {
	gorm.Model
	Description string
	Link_to     string
	Priority    int
	Attachment  Attachment `gorm:"polymorphic:Attachment;"`
}
