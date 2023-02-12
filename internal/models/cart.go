package internal

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID   uint
	User     User
	CartItem []CartItem
}
