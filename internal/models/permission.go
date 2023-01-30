package internal

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	PageName  string
	PageRoute string
}
