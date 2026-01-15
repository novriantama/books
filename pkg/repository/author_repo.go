package repository

import (
	"books/pkg/models"
	"books/pkg/utils"
	"strings"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	CreateAuthor(author *models.Author) error
	UpdateAuthor(author *models.Author) error
	DeleteAuthor(id uint) error
	FindByID(id uint) (*models.Author, error)
	FindAll(pagination utils.Pagination) (*[]models.Author, int64, error)
}

type authorRepo struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepo{db}
}

func (r *authorRepo) CreateAuthor(author *models.Author) error {
	return r.db.Create(author).Error
}

func (r *authorRepo) UpdateAuthor(author *models.Author) error {
	return r.db.Save(author).Error
}

func (r *authorRepo) DeleteAuthor(id uint) error {
	return r.db.Delete(&models.Author{}, id).Error
}

func (r *authorRepo) FindByID(id uint) (*models.Author, error) {
	var author models.Author
	// Preload Books to see what books this author has written
	err := r.db.Preload("Books").First(&author, id).Error
	return &author, err
}

func (r *authorRepo) FindAll(p utils.Pagination) (*[]models.Author, int64, error) {
	var authors []models.Author
	var totalRows int64

	db := r.db.Model(&models.Author{})

	if p.Search != "" {
		search := "%" + strings.ToLower(p.Search) + "%"
		db = db.Where("LOWER(name) LIKE ? OR LOWER(bio) LIKE ?", search, search)
	}

	db.Count(&totalRows)
	err := db.Scopes(p.Paginate()).Preload("Books").Find(&authors).Error
	return &authors, totalRows, err
}
