package internal

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string
}
