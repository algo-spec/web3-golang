package models

import "time"

type User struct {
	ID        uint `gorm:"primary_key"`
	Username  string
	Password  string
	Email     string
	Posts     []Post
	CreatedAt time.Time
	UpdatedAt time.Time
}
