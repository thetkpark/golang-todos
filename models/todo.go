package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title      string
	IsFinished bool
	UserId     uint
}
