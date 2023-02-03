package internal

import "gorm.io/gorm"

type Banner struct {
	gorm.Model
	Description string `json:"description"`
	Link_to     string `json:"link_to"`
	Priority    int    `json:"priority"`
}
