package posts

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	posts, err := c.usecase.GetPosts(limit_int, offset_int)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, posts)
}
