package internal

type CartItem struct {
	Base
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
	CartID    uint    `json:"cart_id"`
	Cart      Cart    `json:"cart"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product"`
	Sku       Sku     `json:"sku"`
}

type CreateCartItemInput struct {
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
	CartID    uint    `json:"cart_id"`
	ProductID uint    `json:"product_id"`
}
