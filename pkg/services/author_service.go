package services

import (
	"books/pkg/models"
	"books/pkg/repository"
	"books/pkg/utils"
	"errors"
)

type AuthorService interface {
	CreateAuthor(author models.Author) error
	GetAuthor(id uint) (*models.Author, error)
	GetAllAuthors(p utils.Pagination) (*[]models.Author, int64, error)
	UpdateAuthor(id uint, author models.Author) error
	DeleteAuthor(id uint) error
}

type authorService struct {
	repo repository.AuthorRepository
}

func NewAuthorService(repo repository.AuthorRepository) AuthorService {
	return &authorService{repo}
}

func (s *authorService) CreateAuthor(author models.Author) error {
	return s.repo.CreateAuthor(&author)
}

func (s *authorService) GetAuthor(id uint) (*models.Author, error) {
	return s.repo.FindByID(id)
}

func (s *authorService) GetAllAuthors(p utils.Pagination) (*[]models.Author, int64, error) {
	return s.repo.FindAll(p)
}

func (s *authorService) UpdateAuthor(id uint, input models.Author) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("author not found")
	}
	existing.Name = input.Name
	existing.Bio = input.Bio
	return s.repo.UpdateAuthor(existing)
}

func (s *authorService) DeleteAuthor(id uint) error {
	return s.repo.DeleteAuthor(id)
}
