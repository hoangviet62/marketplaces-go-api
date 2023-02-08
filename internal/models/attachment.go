package internal

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model
	AttachmentType string `json:"attachment_type"`
	AttachmentID   uint   `json:"attachment_id"`
	Url            string `json:"url"`
}
