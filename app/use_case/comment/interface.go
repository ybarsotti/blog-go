package comment

import "github.com/ybarsotti/blog-test/entity"

type UseCase interface {
	Create(postId int, comment string, author string) (*entity.Comment, error)
}
