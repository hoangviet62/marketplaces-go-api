package internal

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Code       string
	Status     int
	UserID     int32
	User       User
	OrderItems []OrderItem
}
