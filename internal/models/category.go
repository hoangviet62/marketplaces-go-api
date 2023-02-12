package internal

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string
	Products    []Product
	Images      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:CategoryImages"`
	Medias      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:CategoryMedias"`
	Attachments []Attachment `json:"Attachments,omitempty" gorm:"polymorphic:Attachment;"`
}

type CreateCategoryInput struct {
	Name string
}
