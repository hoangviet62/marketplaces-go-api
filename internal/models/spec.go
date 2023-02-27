package internal

type Spec struct {
	Base
	Description string `json:"description"`
	SkuID       uint   `json:"sku_id"`
	ProductID   uint   `json:"product_id"`
}
