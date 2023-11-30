package comment

import "github.com/ybarsotti/blog-test/entity"

type UseCase interface {
	Create(postId int, comment string, author string) (*entity.Comment, error)
	GetAllByPost(postId int) ([]*entity.Comment, error)
	Delete(postId int, commentId int) error
}
