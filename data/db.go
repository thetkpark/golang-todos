package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("todo.data"), &gorm.Config{})
}
