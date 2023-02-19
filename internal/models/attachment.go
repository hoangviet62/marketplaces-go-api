package internal

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model
	AttachmentType string `json:"attachment"`
	AttachmentID   uint `json:"attachment_id"`
	Url            string `json:"url"`
}
