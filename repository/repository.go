package repository

import (
	"gorm.io/gorm"

	irepository "github.com/ybarsotti/blog-test/app/repository"
	"github.com/ybarsotti/blog-test/entity"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) irepository.PostRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(p *entity.Post) (*entity.Post, error) {
	dbPost := FromPostEntity(p)

	result := r.db.Create(&dbPost)

	if result.Error != nil {
		return nil, result.Error
	}

	return dbPost.ToPostEntity(), nil
}

func (r *repository) Update(p *entity.Post) (*entity.Post, error) {
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

func (r *repository) GetAll() ([]*entity.Post, error) {
	var posts []Post
	var entityPosts []*entity.Post
	r.db.Order("id").Find(&posts)

	for _, post := range posts {
		entityPosts = append(entityPosts, post.ToPostEntity())
	}

	return entityPosts, nil
}

func (r *repository) GetByID(id int) *entity.Post {
	var post Post
	r.db.Where("id = ?", id).First(&post)
	return post.ToPostEntity()
}

func (r *repository) DeleteById(id int) {
	r.db.Delete(&Post{}, id)
}
