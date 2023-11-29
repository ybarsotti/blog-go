package repository

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	Content   string
	Author    string
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Comment struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	PostID    string
	Post      Post
	Author    string
	Comment   string
	UpdatedAt time.Time
	CreatedAt time.Time
}
