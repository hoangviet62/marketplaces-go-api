package internal

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Code       string
	Status     int
	UserID     uint
	User       User
	OrderItems []OrderItem
}
