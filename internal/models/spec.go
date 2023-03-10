package internal

import "gorm.io/datatypes"

type Spec struct {
	Base
	Description datatypes.JSON `gorm:"not null" json:"description"`
	ProductID   uint           `json:"product_id"`
}
