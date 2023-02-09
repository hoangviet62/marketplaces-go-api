package internal

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name       string
	Products   []Product
	Attachment Attachment `gorm:"polymorphic:Attachment;"`
}

type CreateCategoryInput struct {
	Name string
}
