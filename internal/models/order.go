package internal

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Code       string
	Status     uint
	UserID     uint
	User       User
	OrderItems []OrderItem
}
