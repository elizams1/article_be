package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"article_be/internal/posts"

	// _ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var postsController *posts.PostsController

func Handler(w http.ResponseWriter, r *http.Request) {
	// Membuat router Gin untuk menangani HTTP request
	router := gin.Default()

	// Mendefinisikan rute
	router.POST("/article", postsController.CreateUser)
	router.GET("/article/list/:limit/:offset", postsController.GetPosts)
	router.GET("/article/:id", postsController.GetPostByID)
	router.PUT("/article/:id", postsController.UpdatePost)
	router.DELETE("/article/:id", postsController.DeletePost)

	// Menangani request
	router.ServeHTTP(w, r)
}
func main() {

	caDoc := "C:\\Users\\USER\\Downloads\\ca.pem"

	if _, err := os.Stat(caDoc); os.IsNotExist(err) {
		log.Fatal("File CA certificate tidak ditemukan di:", caDoc)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("AIVEN_CONNECTION_STRING")
	log.Println(("connectionString: " + connectionString))
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal membuka koneksi:", err)
	}

	err = db.AutoMigrate(&posts.Posts{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi tabel:", err)
	}

	fmt.Println("Berhasil terhubung ke database Aiven MySQL!")

	// Inisialisasi repository dan usecase
	postRepo := posts.NewPostsRepository(db)
	postUsecase := posts.NewPostsUsecase(postRepo)
	postsController := posts.NewPostsController(postUsecase)

	router := gin.Default()

	router.POST("/article", postsController.CreateUser)
	router.GET("/article/list/:limit/:offset", postsController.GetPosts)
	router.GET("/article/:id", postsController.GetPostByID)
	router.PUT("/article/:id", postsController.UpdatePost)
	router.DELETE("/article/:id", postsController.DeletePost)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
