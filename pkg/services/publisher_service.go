package services

import (
	"books/pkg/models"
	"books/pkg/repository"
	"books/pkg/utils"
	"errors"
)

type PublisherService interface {
	CreatePublisher(publisher models.Publisher) error
	GetPublisher(id uint) (*models.Publisher, error)
	GetAllPublishers(p utils.Pagination) (*[]models.Publisher, int64, error)
	UpdatePublisher(id uint, publisher models.Publisher) error
	DeletePublisher(id uint) error
}

type publisherService struct {
	repo repository.PublisherRepository
}

func NewPublisherService(repo repository.PublisherRepository) PublisherService {
	return &publisherService{repo}
}

func (s *publisherService) CreatePublisher(publisher models.Publisher) error {
	return s.repo.CreatePublisher(&publisher)
}

func (s *publisherService) GetPublisher(id uint) (*models.Publisher, error) {
	return s.repo.FindByID(id)
}

func (s *publisherService) GetAllPublishers(p utils.Pagination) (*[]models.Publisher, int64, error) {
	return s.repo.FindAll(p)
}

func (s *publisherService) UpdatePublisher(id uint, input models.Publisher) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("publisher not found")
	}
	existing.Name = input.Name
	existing.Address = input.Address
	return s.repo.UpdatePublisher(existing)
}

func (s *publisherService) DeletePublisher(id uint) error {
	return s.repo.DeletePublisher(id)
}
