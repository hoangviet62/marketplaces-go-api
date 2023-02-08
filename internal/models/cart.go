package internal

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID   uint       `json:"user_id"`
	CartItem []CartItem `json:"cart_items"`
}
