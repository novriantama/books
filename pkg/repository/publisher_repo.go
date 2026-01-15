package repository

import (
	"books/pkg/models"
	"books/pkg/utils"
	"strings"

	"gorm.io/gorm"
)

type PublisherRepository interface {
	CreatePublisher(publisher *models.Publisher) error
	UpdatePublisher(publisher *models.Publisher) error
	DeletePublisher(id uint) error
	FindByID(id uint) (*models.Publisher, error)
	FindAll(pagination utils.Pagination) (*[]models.Publisher, int64, error)
}

type publisherRepo struct {
	db *gorm.DB
}

func NewPublisherRepository(db *gorm.DB) PublisherRepository {
	return &publisherRepo{db}
}

func (r *publisherRepo) CreatePublisher(publisher *models.Publisher) error {
	return r.db.Create(publisher).Error
}

func (r *publisherRepo) UpdatePublisher(publisher *models.Publisher) error {
	return r.db.Save(publisher).Error
}

func (r *publisherRepo) DeletePublisher(id uint) error {
	return r.db.Delete(&models.Publisher{}, id).Error
}

func (r *publisherRepo) FindByID(id uint) (*models.Publisher, error) {
	var publisher models.Publisher
	err := r.db.First(&publisher, id).Error
	return &publisher, err
}

func (r *publisherRepo) FindAll(p utils.Pagination) (*[]models.Publisher, int64, error) {
	var publishers []models.Publisher
	var totalRows int64

	db := r.db.Model(&models.Publisher{})

	if p.Search != "" {
		search := "%" + strings.ToLower(p.Search) + "%"
		db = db.Where("LOWER(name) LIKE ?", search)
	}

	db.Count(&totalRows)
	err := db.Scopes(p.Paginate()).Find(&publishers).Error
	return &publishers, totalRows, err
}
