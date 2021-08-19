package models

import "time"

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `gorm:"unique"`
	Password  string
	Todos     []Todo `gorm:"foreignKey:UserId"`
}
