package repository

import (
	"github.com/ybarsotti/blog-test/entity"
)

func (c *Comment) ToCommentEntity() *entity.Comment {
	comment := &entity.Comment{
		ID:        c.ID,
		Author:    c.Author,
		Comment:   c.Comment,
		UpdatedAt: c.UpdatedAt,
		CreatedAt: c.CreatedAt,
	}

	if c.Post != nil {
		comment.Post = c.Post.ToPostEntity()
	}

	return comment
}

func FromCommentEntity(e *entity.Comment) *Comment {
	return &Comment{
		ID:        e.ID,
		Post:      FromPostEntity(e.Post),
		Author:    e.Author,
		Comment:   e.Comment,
		UpdatedAt: e.UpdatedAt,
		CreatedAt: e.CreatedAt,
	}
}
