package internal

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Code       string      `json:"code"`
	Status     int         `json:"status"`
	UserID     uint        `json:"user_id"`
	User       User        `json:"user"`
	OrderItems []OrderItem `json:"order_items"`
}
