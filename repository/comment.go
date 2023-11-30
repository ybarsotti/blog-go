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

func (r repository) GetAllByPost(postId int) []*entity.Comment {
	var comments []*entity.Comment
	var dbComments []*Comment
	r.db.Where("post_id = ?", postId).Find(&dbComments)

	for _, dbComment := range dbComments {
		comments = append(comments, dbComment.ToCommentEntity())
	}

	return comments
}

func (r repository) DeleteByID(id int) {
	r.db.Delete(&Comment{}, id)
}
