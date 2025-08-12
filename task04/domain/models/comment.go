package models

import "time"

type Comment struct {
	ID        uint `gorm:"primary_key"`
	Content   string
	PostID    uint
	UserID    uint
	Post      Post
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}
