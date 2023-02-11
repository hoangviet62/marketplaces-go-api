package internal

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID   int32
	User     User
	CartItem []CartItem
}
