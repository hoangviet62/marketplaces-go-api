package internal

type Order struct {
	Base
	Code       string      `json:"code"`
	Status     uint        `json:"status"`
	UserID     uint        `json:"user_id"`
	User       User        `json:"user"`
	OrderItems []OrderItem `json:"order_items,omitempty"`
}
