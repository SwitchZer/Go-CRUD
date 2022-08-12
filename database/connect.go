package database

import (
	"blogbackend/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error file .env")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Tidak bisa Connect ke database")
	} else {
		log.Println("Koneksi Berhasil")
	}
	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}
