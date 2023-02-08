package internal

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name   string `json:"name"`
	Status int    `json:"status"`
	Code   string `json:"code"`
}
