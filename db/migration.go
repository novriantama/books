package db

import (
	"books/pkg/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Author{}, &models.Publisher{}, &models.Book{}, &models.TokenBlacklist{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Database migration completed.")
}
