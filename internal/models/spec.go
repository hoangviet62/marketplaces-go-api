package internal

import "gorm.io/gorm"

type Spec struct {
	gorm.Model
	Description string
	SkuID       uint
	ProductID   uint
}
