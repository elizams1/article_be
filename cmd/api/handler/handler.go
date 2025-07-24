package handler

import (
	"article_be/posts"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	postsController *posts.PostsController
)

func Handler(w http.ResponseWriter, r *http.Request) {
	connectionString := os.Getenv("AIVEN_CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err = db.AutoMigrate(&posts.Posts{}); err != nil {
		log.Fatal("Failed to migrate tables:", err)
	}

	log.Println("Connected to the database")

	postRepo := posts.NewPostsRepository(db)
	postUsecase := posts.NewPostsUsecase(postRepo)
	postsController = posts.NewPostsController(postUsecase)

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"} // Allow all origins (you can restrict this)
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true

	router.Use(cors.New(corsConfig))

	router.POST("/article", postsController.CreateUser)
	router.GET("/article/list/:limit/:offset", postsController.GetPosts)
	router.GET("/article/:id", postsController.GetPostByID)
	router.PUT("/article/:id", postsController.UpdatePost)
	router.DELETE("/article/:id", postsController.DeletePost)

	// Serve HTTP request
	router.ServeHTTP(w, r)
}
