package repository

import (
	"gorm.io/gorm"

	irepository "github.com/ybarsotti/blog-test/app/repository"
	"github.com/ybarsotti/blog-test/entity"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) irepository.PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(p *entity.Post) (*entity.Post, error) {
	dbPost := FromPostEntity(p)

	result := r.db.Create(&dbPost)

	if result.Error != nil {
		return nil, result.Error
	}

	return dbPost.ToPostEntity(), nil
}

func (r *postRepository) Update(p *entity.Post) (*entity.Post, error) {
	updateData := &Post{
		Title:   p.Title,
		Content: p.Content,
	}
	result := r.db.Model(&Post{}).Where("id = ?", p.ID).Updates(updateData)

	if result.Error != nil {
		return nil, result.Error
	}

	var updatedPost Post
	r.db.First(&updatedPost, p.ID)

	return updatedPost.ToPostEntity(), nil
}

func (r *postRepository) GetAll() ([]*entity.Post, error) {
	var posts []Post
	var entityPosts []*entity.Post
	r.db.Order("id").Find(&posts)

	for _, post := range posts {
		entityPosts = append(entityPosts, post.ToPostEntity())
	}

	return entityPosts, nil
}

func (r *postRepository) GetByID(id int) *entity.Post {
	var post Post
	r.db.Where("id = ?", id).First(&post)
	return post.ToPostEntity()
}

func (r *postRepository) DeleteById(id int) error {
	tx := r.db.Begin()
	if err := tx.Unscoped().Where("post_id = ?", id).Delete(&Comment{}).Error; err != nil {
		tx.Rollback()
		return nil
	}

	if err := tx.Unscoped().Delete(&Post{}, id).Error; err != nil {
		tx.Rollback()
		return nil
	}

	tx.Commit()
	return r.db.Error
}
