package internal

type Cart struct {
	Base
	UserID   uint       `json:"user_id"`
	User     User       `json:"user"`
	CartItem []CartItem `json:"cart_items"`
}
