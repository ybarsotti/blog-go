package repository

import "github.com/ybarsotti/blog-test/entity"

type CommentRepository interface {
	Create(c *entity.Comment) (*entity.Comment, error)
	//GetAllByPost()
	//Delete()
}
