package internal

import (
	"database/sql/driver"
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
	Base
	Username  string   `json:"username" gorm:"size:255;index:idx_name,unique"`
	Password  string   `json:"-" gorm:"size:255"`
	Email     string   `json:"email" gorm:"size:255;index:idx_email,unique"`
	Mobile    string   `json:"mobile" gorm:"size:255;index:idx_mobile,unique"`
	Status    Status   `json:"status" sql:"type:ENUM('inactive', 'active')"`
	Country   *Country `json:"country"`
	CountryId *uint    `json:"country_id,omitempty"`
	Role      Role     `json:"role" sql:"type:ENUM('admin', 'customer', 'seller')" gorm:"default:customer"` // MySQL
}

type SignUpInput struct {
	Username        string `binding:"required"`
	Email           string `binding:"required"`
	Password        string `binding:"required,min=8"`
	PasswordConfirm string `binding:"required"`
	Mobile          string `binding:"required"`
	Role            Role   `sql:"type:ENUM('admin', 'customer', 'seller')" gorm:"default:customer"`
}

type SignInInput struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
