package internal

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model
	Attachment_type string `json:"attachment_type"`
	Attachment_id   int    `json:"attachment_id"`
	Url             string `json:"url"`
}
