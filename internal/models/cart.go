package internal

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID   uint       `json:"user_id"`
	User     User       `json:"user"`
	CartItem []CartItem `json:"cart_items"`
}
