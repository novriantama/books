package db

import (
	"books/pkg/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := models.User{
		Email:    "admin@example.com",
		Password: string(hashedPwd),
	}

	if err := db.FirstOrCreate(&user, models.User{Email: user.Email}).Error; err != nil {
		log.Println("Seeding failed:", err)
	} else {
		log.Println("Database seeded with admin user.")
	}
}
