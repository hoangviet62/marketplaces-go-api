package internal

type Category struct {
	Base
	Name        string       `json:"name"`
	Products    []Product    `gorm:"constraint:OnUpdate:RESTRICT;" json:"products,omitempty"`
	Images      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:CategoryImages" json:"images,omitempty"`
	Medias      []Attachment `gorm:"polymorphic:Attachment;polymorphicValue:CategoryMedias" json:"medias,omitempty"`
	Attachments []Attachment `gorm:"polymorphic:Attachment;" json:"attachments,omitempty"`
}

type CreateCategoryInput struct {
	Name string
}
