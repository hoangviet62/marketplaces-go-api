package internal

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type RoleType string

const (
	Admin    RoleType = "admin"
	Seller   RoleType = "seller"
	Customer RoleType = "customer"
)

func (e *RoleType) Scan(value interface{}) error {
	*e = RoleType(value.([]byte))
	return nil
}

func (e RoleType) Value() (driver.Value, error) {
	return string(e), nil
}

type Role struct {
	gorm.Model
	RoleType RoleType
}
