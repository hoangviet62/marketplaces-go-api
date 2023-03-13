package internal

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Cart struct {
	Base
	UserID   uint       `json:"user_id"`
	CartItem []CartItem `gorm:"constraint:OnDelete:CASCADE;" json:"cart_items"`
}

type CreateCartParams struct {
	Product Product
	Cart    Cart
}

func (cart *Cart) AfterCreate(tx *gorm.DB) (err error) {
	tx.Preload(clause.Associations).Preload("CartItem.Product.Images").Preload("CartItem.Product.Medias").First(&cart, "id = ?", cart.ID)
	return
}
