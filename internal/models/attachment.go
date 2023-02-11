package internal

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model
	AttachmentType string
	AttachmentID   int32
	Url            string
}
