package internal

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Code   string `json:"code"`
	Status int    `json:"status"`
	UserID int
	User   User
}
