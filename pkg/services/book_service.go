package services

import (
	"books/pkg/models"
	"books/pkg/repository"
	"books/pkg/utils"
	"errors"
)

type BookService interface {
	CreateBook(book models.Book) error
	GetBook(id uint) (*models.Book, error)
	GetAllBooks(p utils.Pagination) (*[]models.Book, int64, error)
	UpdateBook(id uint, book models.Book) error
	DeleteBook(id uint) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) CreateBook(book models.Book) error {
	return s.repo.CreateBook(&book)
}

func (s *bookService) GetBook(id uint) (*models.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) GetAllBooks(p utils.Pagination) (*[]models.Book, int64, error) {
	return s.repo.FindAll(p)
}

func (s *bookService) UpdateBook(id uint, input models.Book) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("book not found")
	}

	existing.Title = input.Title
	existing.Description = input.Description
	existing.AuthorID = input.AuthorID
	existing.PublisherID = input.PublisherID

	return s.repo.UpdateBook(existing)
}

func (s *bookService) DeleteBook(id uint) error {
	return s.repo.DeleteBook(id)
}
