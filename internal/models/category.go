package internal

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Products    []Product `json:"products,omitempty"`
	Images      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:CategoryImages" json:"images,omitempty"`
	Medias      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:CategoryMedias" json:"medias,omitempty"`
	Attachments []Attachment `gorm:"polymorphic:Attachment;" json:"attachments,omitempty"`
}

type CreateCategoryInput struct {
	Name string
}
