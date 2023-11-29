package repository

import (
	"github.com/ybarsotti/blog-test/entity"
)

func (p *Post) ToPostEntity() *entity.Post {
	return &entity.Post{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		Author:    p.Author,
		UpdatedAt: p.UpdatedAt,
		CreatedAt: p.CreatedAt,
	}
}

func FromPostEntity(e *entity.Post) *Post {
	return &Post{
		ID:        e.ID,
		Title:     e.Title,
		Content:   e.Content,
		Author:    e.Author,
		UpdatedAt: e.UpdatedAt,
		CreatedAt: e.CreatedAt,
	}
}
