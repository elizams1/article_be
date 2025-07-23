package main

import (
	"fmt"
	"log"
	"os"

	"article_be/internal/posts"

	// _ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	caDoc := "C:\\Users\\USER\\Downloads\\ca.pem"

	if _, err := os.Stat(caDoc); os.IsNotExist(err) {
		log.Fatal("File CA certificate tidak ditemukan di:", caDoc)
	}

	connectionString := os.Getenv("AIVEN_CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal membuka koneksi:", err)
	}

	err = db.AutoMigrate(&posts.Posts{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi tabel:", err)
	}

	fmt.Println("Berhasil terhubung ke database Aiven MySQL!")

}
