package internal

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model
	Attachment_type string `json:"attachment_type"`
	AttachmentID    uint   `json:"attachment_id"`
	Banner          Banner `gorm:"foreignKey:AttachmentID"`
	// Category        Category `gorm:"foreignKey:AttachmentID"`
	// Product         Product  `gorm:"foreignKey:AttachmentID"`
	// Sku             Sku      `gorm:"foreignKey:AttachmentID"`
	Url string `json:"url"`
}
