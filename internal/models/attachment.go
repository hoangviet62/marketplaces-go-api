package internal

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model
	AttachmentType string
	AttachmentID   uint
	Url            string
}
