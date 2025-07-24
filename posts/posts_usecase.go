package posts

type PostsUsecase interface {
	GetPosts(limit int, offset int) ([]*Posts, error)
	GetPostByID(id int) (*Posts, error)
	CreatePost(post *Posts) error
	UpdatePost(id int, post *Posts) error
	DeletePost(id int) error
	SoftDeletePost(id int) error
}

type postsUsecase struct {
	repository PostsRepository
}

func NewPostsUsecase(repository PostsRepository) PostsUsecase {
	return &postsUsecase{repository}
}

func (u *postsUsecase) GetPosts(limit int, offset int) ([]*Posts, error) {
	return u.repository.GetPosts(limit, offset)
}
func (u *postsUsecase) GetPostByID(id int) (*Posts, error) {
	return u.repository.GetPostByID(id)
}
func (u *postsUsecase) CreatePost(post *Posts) error {
	return u.repository.CreatePost(post)
}
func (u *postsUsecase) UpdatePost(id int, post *Posts) error {
	return u.repository.UpdatePost(id, post)
}
func (u *postsUsecase) DeletePost(id int) error {
	return u.repository.DeletePost(id)
}

func (u *postsUsecase) SoftDeletePost(id int) error {
	return u.repository.SoftDeletePost(id)
}
