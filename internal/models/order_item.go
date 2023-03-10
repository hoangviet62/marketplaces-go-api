package internal

type OrderItem struct {
	Base
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	SkuID     *uint   `json:"sku_id,omitempty"`
}
