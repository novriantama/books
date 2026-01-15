package db

import (
	"books/pkg/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// 1. Seed User
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

	// 2. Seed Authors
	authors := []models.Author{
		{Name: "J.K. Rowling", Bio: "British author, best known for the Harry Potter series."},
		{Name: "George R.R. Martin", Bio: "American novelist and short story writer, creating A Song of Ice and Fire."},
		{Name: "J.R.R. Tolkien", Bio: "English writer, poet, philologist, and academic."},
	}

	for i, author := range authors {
		// We use FirstOrCreate to ensure we don't duplicate data on restart
		if err := db.FirstOrCreate(&authors[i], models.Author{Name: author.Name}).Error; err != nil {
			log.Println("Failed to seed author:", err)
		}
	}

	// 3. Seed Publishers
	publishers := []models.Publisher{
		{Name: "Bloomsbury", Address: "London, UK"},
		{Name: "Bantam Books", Address: "New York, USA"},
		{Name: "Allen & Unwin", Address: "Sydney, Australia"},
	}

	for i, publisher := range publishers {
		if err := db.FirstOrCreate(&publishers[i], models.Publisher{Name: publisher.Name}).Error; err != nil {
			log.Println("Failed to seed publisher:", err)
		}
	}

	// 4. Seed Books
	books := []models.Book{
		{
			Title:       "Harry Potter and the Philosopher's Stone",
			Description: "A young wizard discovers his magical heritage.",
			AuthorID:    authors[0].ID,    // J.K. Rowling
			PublisherID: publishers[0].ID, // Bloomsbury
		},
		{
			Title:       "A Game of Thrones",
			Description: "Noble families fight for control of the Iron Throne.",
			AuthorID:    authors[1].ID,    // George R.R. Martin
			PublisherID: publishers[1].ID, // Bantam Books
		},
		{
			Title:       "The Hobbit",
			Description: "Bilbo Baggins goes on an adventure.",
			AuthorID:    authors[2].ID,    // J.R.R. Tolkien
			PublisherID: publishers[2].ID, // Allen & Unwin
		},
	}

	for _, book := range books {
		if err := db.FirstOrCreate(&book, models.Book{Title: book.Title}).Error; err != nil {
			log.Println("Failed to seed book:", err)
		}
	}

	log.Println("Database seeding completed successfully.")
}
