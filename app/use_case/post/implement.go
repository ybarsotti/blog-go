package post

import (
	"github.com/ybarsotti/blog-test/app/repository"
	"github.com/ybarsotti/blog-test/entity"
)

type usecase struct {
	post_repo repository.PostRepository
}

func NewUseCase(post_repo repository.PostRepository) UseCase {
	return &usecase{post_repo: post_repo}
}

func (uc *usecase) Create(title string, content string, author string) (*entity.Post, error) {
	post := entity.Post{
		Title:   title,
		Content: content,
		Author:  author,
	}
	p, err := uc.post_repo.Create(&post)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (uc *usecase) GetAll() ([]*entity.Post, error) {
	posts, err := uc.post_repo.GetAll()

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (uc *usecase) Get(id int) *entity.Post {
	return uc.post_repo.GetByID(id)
}

func (uc *usecase) Update(id int, title string, content string) (*entity.Post, error) {
	post := entity.Post{
		ID:      id,
		Title:   title,
		Content: content,
	}

	p, err := uc.post_repo.Update(&post)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (uc *usecase) Delete(id int) error {
	return uc.post_repo.DeleteById(id)
}
