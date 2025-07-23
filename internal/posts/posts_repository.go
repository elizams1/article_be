package posts

import "gorm.io/gorm"

type PostsRepository interface {
	GetPosts(limit int, offset int) ([]*Posts, error)
	CreatePost(post *Posts) error
}

type postsRepository struct {
	db *gorm.DB
}

func NewPostsRepository(db *gorm.DB) PostsRepository {
	return &postsRepository{db}
}

func (r *postsRepository) GetPosts(limit int, offset int) ([]*Posts, error) {
	var posts []*Posts
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postsRepository) CreatePost(post *Posts) error {
	query := `INSERT INTO posts (title, content, category, status) VALUES (?, ?, ?, ?) RETURNING id`
	res := r.db.Exec(query, post.Title, post.Content, post.Category, post.Status).Scan(&post.ID)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
