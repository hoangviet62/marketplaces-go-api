package internal

import "database/sql/driver"

type OrderStatus string

const (
	ApprovedOrder   OrderStatus = "approved"
	ProcessingOrder OrderStatus = "processing"
	RejectedOrder   OrderStatus = "rejected"
)

func (r *OrderStatus) Scan(value interface{}) error {
	*r = OrderStatus(value.([]byte))
	return nil
}

func (r OrderStatus) Value() (driver.Value, error) {
	return string(r), nil
}

type Order struct {
	Base
	Code       string      `json:"code,omitempty"`
	Status     string      `json:"status" gorm:"default:processing"`
	UserID     uint        `json:"user_id"`
	OrderItems []OrderItem `json:"order_items,omitempty"`
}

// status default Processing, Approved, Reject
