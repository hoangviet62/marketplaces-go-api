package internal

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type Status string

const (
	Inactive Status = "inactive"
	Acttive  Status = "active"
)

func (e *Status) Scan(value interface{}) error {
	*e = Status(value.([]byte))
	return nil
}

func (e Status) Value() (driver.Value, error) {
	return string(e), nil
}

type Role string

const (
	Admin    Role = "admin"
	Customer Role = "customer"
	Seller   Role = "seller"
)

func (r *Role) Scan(value interface{}) error {
	*r = Role(value.([]byte))
	return nil
}

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}

type User struct {
	gorm.Model
	Username  string `gorm:"size:255;index:idx_name,unique"`
	Password  string `gorm:"size:255" json:",omitempty"`
	Email     string `gorm:"size:255;index:idx_email,unique"`
	Mobile    string `gorm:"size:255;index:idx_mobile,unique"`
	Status    Status `json:"Status" sql:"type:ENUM('inactive', 'active')"`
	Country   *Country
	CountryId *uint `json:"Country,omitempty"`
	Role      Role  `json:"role" sql:"type:ENUM('admin', 'customer', 'seller')"` // MySQL
}

type SignUpInput struct {
	Username        string `binding:"required"`
	Email           string `binding:"required"`
	Password        string `binding:"required,min=8"`
	PasswordConfirm string `binding:"required"`
	Mobile          string `binding:"required"`
}

type SignInInput struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
