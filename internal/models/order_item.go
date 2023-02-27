package internal

type OrderItem struct {
	Base
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
	OrderID  uint    `json:"order_id"`
	Order    Order   `json:"order"`
	Product  Product `json:"product"`
	Sku      Sku     `json:"sku"`
}
