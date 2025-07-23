package posts

import (
	"time"

	"gorm.io/gorm"
)

type PostsRepository interface {
	GetPosts(limit int, offset int) ([]*Posts, error)
	GetPostByID(id int) (*Posts, error)
	CreatePost(post *Posts) error
	UpdatePost(id int, post *Posts) error
	DeletePost(id int) error
}

type postsRepository struct {
	db *gorm.DB
}

func NewPostsRepository(db *gorm.DB) PostsRepository {
	return &postsRepository{db}
}

func (r *postsRepository) GetPosts(limit int, offset int) ([]*Posts, error) {
	var posts []*Posts
	if err := r.db.Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postsRepository) GetPostByID(id int) (*Posts, error) {
	var post Posts
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postsRepository) CreatePost(post *Posts) error {
	return r.db.Create(post).Error
}

func (r *postsRepository) UpdatePost(id int, post *Posts) error {
	var existingPost Posts
	err := r.db.First(&existingPost, id)
	if err != nil {
		return err.Error
	}
	post.ID = existingPost.ID
	if post.Title == "" {
		post.Title = existingPost.Title
	}
	if post.Content == "" {
		post.Content = existingPost.Content
	}
	if post.Category == "" {
		post.Category = existingPost.Category
	}
	if post.Status == "" {
		post.Status = existingPost.Status
	}
	post.CreatedDate = existingPost.CreatedDate
	post.UpdatedDate = time.Now().UTC()

	return r.db.Save(&post).Error
}

func (r *postsRepository) DeletePost(id int) error {
	var existingPost Posts
	err := r.db.First(&existingPost, id)
	if err != nil {
		return err.Error
	}
	return r.db.Delete(&existingPost, id).Error
}
