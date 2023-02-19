package internal

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name   string `json:"name"`
	Status uint `json:"status"`
	Code   string `json:"code"`
}
