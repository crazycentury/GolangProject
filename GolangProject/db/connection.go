package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	var err error

	// Konfigurasi koneksi database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),     // Contoh: "localhost"
		os.Getenv("DB_USER"),     // Contoh: "postgres"
		os.Getenv("DB_PASSWORD"), // Contoh: "password"
		os.Getenv("DB_NAME"),     // Contoh: "myapp_db"
		os.Getenv("DB_PORT"),     // Contoh: "5432"
	)

	// Koneksi ke PostgreSQL menggunakan GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
}
