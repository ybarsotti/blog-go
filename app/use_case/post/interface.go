package post

import "github.com/ybarsotti/blog-test/entity"

type UseCase interface {
	Create(title string, content string, author string) (*entity.Post, error)
	GetAll() ([]*entity.Post, error)
	Get(id int) *entity.Post
	Update(id int, title string, content string) (*entity.Post, error)
	Delete(id int) error
}
