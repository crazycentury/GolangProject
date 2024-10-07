package migrations

import (
	"GolangProject/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Transaction{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}
