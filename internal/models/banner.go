package internal

import "gorm.io/gorm"

type Banner struct {
	gorm.Model
	Description string `json:"description"`
	LinkTo      string `json:"link_to"`
	Priority    uint` json:"priority"`
	Images      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:banner_images" json:"images,omitempty"`
	Medias      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:category_medias" json:"medias,omitempty"`
	Attachments []Attachment `gorm:"polymorphic:Attachment;" json:"attachments,omitempty"`
}
