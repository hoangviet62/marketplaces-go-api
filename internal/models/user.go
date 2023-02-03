package internal

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Status   int    `json:"status"`
}

type CreateUserInput struct {
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	Email     string `json:"Email"`
	Mobile    string `json:"Mobile"`
	CountryId int    `json:"CountryId"`
	RoleId    int    `json:"RoleId"`
	Status    int    `json:"Status"`
}
