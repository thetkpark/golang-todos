package models

import (
	"time"
)

type Todo struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Title      string    `json:"title"`
	IsFinished bool      `json:"is_finished"`
	UserId     uint      `json:"user_id"`
}
