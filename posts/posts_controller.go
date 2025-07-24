package posts

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type PostsController struct {
	usecase PostsUsecase
}

func NewPostsController(usecase PostsUsecase) *PostsController {
	return &PostsController{usecase: usecase}
}

func (c *PostsController) CreateUser(ctx *gin.Context) {
	var post Posts

	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(post); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if post.Status == "" {
		post.Status = "draft"
	}

	post.CreatedDate = time.Now().UTC()
	post.UpdatedDate = time.Now().UTC()

	if err := c.usecase.CreatePost(&post); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, post)
}

func (c *PostsController) GetPosts(ctx *gin.Context) {
	limit := ctx.Param("limit")
	limit_int, err := strconv.Atoi(limit)

	offset := ctx.Param("offset")
	offset_int, err := strconv.Atoi(offset)

	if limit_int == 0 {
		limit_int = 10
	}
	posts, err := c.usecase.GetPosts(limit_int, offset_int)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, posts)
}

func (c *PostsController) GetPostByID(ctx *gin.Context) {
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	post, err := c.usecase.GetPostByID(id_int)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, post)
}

func (c *PostsController) UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var post Posts
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(post); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.usecase.UpdatePost(id_int, &post); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, post)
}

func (c *PostsController) DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.usecase.DeletePost(id_int); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Post deleted successfully"})
}

func (c *PostsController) SoftDeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.usecase.SoftDeletePost(id_int); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Post soft deleted successfully"})
}
