package repository

import (
	"books/pkg/models"
	"books/pkg/utils"
	"strings"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
	FindByID(id uint) (*models.Book, error)
	FindAll(pagination utils.Pagination) (*[]models.Book, int64, error)
}

type bookRepo struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepo{db}
}

func (r *bookRepo) CreateBook(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepo) UpdateBook(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepo) DeleteBook(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}

func (r *bookRepo) FindByID(id uint) (*models.Book, error) {
	var book models.Book
	// Preload loads the related Author and Publisher data
	err := r.db.Preload("Author").Preload("Publisher").First(&book, id).Error
	return &book, err
}

func (r *bookRepo) FindAll(p utils.Pagination) (*[]models.Book, int64, error) {
	var books []models.Book
	var totalRows int64

	db := r.db.Model(&models.Book{})

	// Apply Search Filter (Title or Description)
	if p.Search != "" {
		search := "%" + strings.ToLower(p.Search) + "%"
		db = db.Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ?", search, search)
	}

	db.Count(&totalRows) // Count total before pagination

	// Apply Pagination & Preload
	err := db.Scopes(p.Paginate()).Preload("Author").Preload("Publisher").Find(&books).Error
	return &books, totalRows, err
}
