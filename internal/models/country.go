package internal

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name   string
	Status int
	Code   string
}
