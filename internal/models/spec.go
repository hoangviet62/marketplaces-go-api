package internal

import "gorm.io/gorm"

type Spec struct {
	gorm.Model
	Description string `json:"description"`
}
