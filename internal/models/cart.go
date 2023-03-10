package internal

type Cart struct {
	Base
	UserID   uint       `json:"user_id"`
	CartItem []CartItem `gorm:"constraint:OnDelete:SET NULL;" json:"cart_items"`
}

type CreateCartParams struct {
	Product Product
	Cart    Cart
}
