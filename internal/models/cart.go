package internal

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID int
	User   User
}
