package internal

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Password  string
	Email     string
	Mobile    string
	CountryId int
	RoleId    int
	Status    int
}
