package internal

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name       string     `json:"name"`
	Products   []Product  `json:"products"`
	Attachment Attachment `gorm:"polymorphic:Attachment;"`
}

type CreateCategoryInput struct {
	Name string `json:"name"`
}
