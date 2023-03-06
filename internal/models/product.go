package internal

import "database/sql/driver"

type ProductStatus string

const (
	Approved ProductStatus = "approved"
	Pending  ProductStatus = "pending"
	Rejected ProductStatus = "rejected"
)

func (r *ProductStatus) Scan(value interface{}) error {
	*r = ProductStatus(value.([]byte))
	return nil
}

func (r ProductStatus) Value() (driver.Value, error) {
	return string(r), nil
}

type Product struct {
	Base
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Tag         string        `json:"tag"`
	CategoryID  uint          `json:"category_id"`
	Spec        Spec          `json:"spec"`
	Skus        []Sku         `json:"skus,omitempty"`
	CartItemID  *uint         `json:"cart_item,omitempty"`
	OrderItemID *uint         `json:"order_item,omitempty"`
	IsFeatured  bool          `json:"is_featured" gorm:"default:false"`
	Status      ProductStatus `json:"status" sql:"type:ENUM('approved', 'pending', 'rejected')" gorm:"default:pending"` // MySQL
	Images      []Attachment  `json:"images,omitempty" gorm:"polymorphic:Attachment;polymorphicValue:product_images"`
	Medias      []Attachment  `json:"medias,omitempty" gorm:"polymorphic:Attachment;polymorphicValue:product_medias"`
	Attachments []Attachment  `json:"attachments,omitempty" gorm:"polymorphic:Attachment;"`
}

// Product => 1 Sku (For now)
// QUERY Product if not admin => only get products status = approved
// Product Status => Default: Pending, 3 ENUMS Pending, Approved, Rejected
// Specs
// path menu: /catalog?type=category||product&id&name
