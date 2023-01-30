package internal

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	FullName  string
	Email     string
	RoleId    int16
	Status    int8
}
