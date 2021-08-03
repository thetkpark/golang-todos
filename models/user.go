package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string
	Password string
	Todos    []Todo `gorm:"foreignKey:UserId"`
}
