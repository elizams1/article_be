package posts

type PostsUsecase interface {
	GetPosts(limit int, offset int) ([]*Posts, error)
	CreatePost(post *Posts) error
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
func (u *postsUsecase) CreatePost(post *Posts) error {
	return u.repository.CreatePost(post)
}
