package comment

import (
	"errors"
	"github.com/ybarsotti/blog-test/app/repository"
	"github.com/ybarsotti/blog-test/entity"
)

type usecase struct {
	comment_repo repository.CommentRepository
	post_repo    repository.PostRepository
}

func NewUseCase(comment_repo repository.CommentRepository, post_repo repository.PostRepository) UseCase {
	return &usecase{comment_repo: comment_repo, post_repo: post_repo}
}

func (u usecase) Create(postId int, comment string, author string) (*entity.Comment, error) {
	post := u.post_repo.GetByID(postId)
	if post.ID == 0 {
		return nil, errors.New("post not found")
	}

	commentE := &entity.Comment{
		Post:    post,
		Author:  author,
		Comment: comment,
	}

	createdComment, err := u.comment_repo.Create(commentE)

	if err != nil {
		return nil, err
	}

	return createdComment, nil
}

func (u usecase) GetAllByPost(postId int) ([]*entity.Comment, error) {
	post := u.post_repo.GetByID(postId)
	if post.ID == 0 {
		return nil, errors.New("post not found")
	}

	comments := u.comment_repo.GetAllByPost(postId)
	return comments, nil
}

func (u usecase) Delete(postId int, commentId int) error {
	post := u.post_repo.GetByID(postId)

	if post.ID == 0 {
		return errors.New("post not found")
	}

	u.comment_repo.DeleteByID(commentId)
	return nil
}
