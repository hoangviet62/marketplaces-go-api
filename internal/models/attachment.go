package internal

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model
	Attachment_type string `json:"attachment_type"`
	AttachmentID    int    `json:"attachment_id"`
	Banner          Banner `gorm:"foreignKey:AttachmentID"`
	Url             string `json:"url"`
}
