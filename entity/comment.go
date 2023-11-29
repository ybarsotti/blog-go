package entity

import "time"

type Comment struct {
	ID        int `json:"id"`
	Post      Post
	Author    string `json:"author"`
	Comment   string `json:"comment"`
	UpdatedAt time.Time
	CreatedAt time.Time `json:"created_at"`
}
