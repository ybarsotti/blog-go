package repository

import "github.com/ybarsotti/blog-test/entity"

type PostRepository interface {
	Create(p *entity.Post) (*entity.Post, error)
	Update(p *entity.Post) (*entity.Post, error)
	GetAll() ([]*entity.Post, error)
	GetByID(id int) *entity.Post
	DeleteById(id int) error
}
