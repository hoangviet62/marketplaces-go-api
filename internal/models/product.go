package internal

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
}
