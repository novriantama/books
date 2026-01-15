package models

import "time"

type Author struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	Books     []Book    `gorm:"foreignKey:AuthorID" json:"books,omitempty"` // 1 Author -> Many Books
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Publisher struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Book struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`

	// Foreign Keys
	AuthorID    uint `json:"author_id"`
	PublisherID uint `json:"publisher_id"`

	// Relationships (for preloading data)
	Author    *Author    `json:"author,omitempty"`
	Publisher *Publisher `json:"publisher,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
