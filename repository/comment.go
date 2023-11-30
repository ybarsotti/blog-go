package repository

import (
	irepository "github.com/ybarsotti/blog-test/app/repository"
	"github.com/ybarsotti/blog-test/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) irepository.CommentRepository {
	return &repository{
		db: db,
	}
}

func (r repository) Create(c *entity.Comment) (*entity.Comment, error) {
	dbComment := FromCommentEntity(c)

	result := r.db.Create(&dbComment)

	if result.Error != nil {
		return nil, result.Error
	}

	return dbComment.ToCommentEntity(), nil
}
