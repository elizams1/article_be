package main

import (
	"log"
	"os"

	"article_be/posts"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	postsController *posts.PostsController
)

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("AIVEN_CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("AIVEN_CONNECTION_STRING is not set")
	}

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err := db.AutoMigrate(&posts.Posts{}); err != nil {
		log.Fatal("Failed to migrate tables:", err)
	}

	log.Println("Connected to the database")

	postRepo := posts.NewPostsRepository(db)
	postUsecase := posts.NewPostsUsecase(postRepo)
	postsController = posts.NewPostsController(postUsecase)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/article", postsController.CreateUser)
	router.GET("/article/list/:limit/:offset", postsController.GetPosts)
	router.GET("/article/:id", postsController.GetPostByID)
	router.PUT("/article/:id", postsController.UpdatePost)
	router.DELETE("/article/:id", postsController.DeletePost)
	return router
}

// Main function for local development
func main() {
	initDB()

	router := setupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}
