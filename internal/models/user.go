package internal

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Password  string
	Email     string
	Mobile    string
	Status    int
	CountryID uint
	Country   Country
	RoleID    uint
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

type UserResponse struct {
	gorm.Model
	Username string `omitempty"`
	Email    string `omitempty"`
	Role     string `omitempty"`
}
