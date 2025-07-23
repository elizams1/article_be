package main

import (
	"article_be/posts"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB
	postsController *posts.PostsController
)

// Fungsi utama yang akan diekspor untuk Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	initDB()
	router := setupRouter()
	router.ServeHTTP(w, r)
}

// Untuk development lokal
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

func initDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("AIVEN_CONNECTION_STRING")

	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal membuka koneksi:", err)
	}

	if err = db.AutoMigrate(&posts.Posts{}); err != nil {
		log.Fatal("Gagal melakukan migrasi tabel:", err)
	}

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
