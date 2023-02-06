package internal

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name   string `json:"name"`
	Status int    `json:"status"`
	UserID uint   `json:"user_id"`
}
