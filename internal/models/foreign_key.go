package internal

import "gorm.io/gorm"

var DB *gorm.DB

func CreateForeignKeys() {
	// DB.Migrator().CreateConstraint(&User{}, "name_checker")
	DB.AutoMigrate()
}
