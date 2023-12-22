package repository

import (
	"gorm.io/gorm/clause"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	Content   string
	Author    string
	Comments  []Comment `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (p *Post) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("post_id = ?", p.ID).Delete(&Comment{})
	return
}

type Comment struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	PostID    int
	Post      *Post
	Author    string
	Comment   string
	UpdatedAt time.Time
	CreatedAt time.Time
}
