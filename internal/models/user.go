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

type User struct {
	gorm.Model
	Username  string `gorm:"size:255;index:idx_name,unique"`
	Password  string `gorm:"size:255" json:",omitempty"`
	Email     string `gorm:"size:255;index:idx_email,unique"`
	Mobile    string `gorm:"size:255;index:idx_mobile,unique"`
	Status    Status `json:"Status" sql:"type:ENUM('inactive', 'active')"`
	CountryId *uint  `json:",omitempty"`
	Country   Country
	RoleId    uint `json:",omitempty"`
	Role      Role
}

type SignUpInput struct {
	Username        string `binding:"required"`
	Email           string `binding:"required"`
	Password        string `binding:"required,min=8"`
	PasswordConfirm string `binding:"required"`
	Mobile          string `binding:"required"`
}

type SignInInput struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
