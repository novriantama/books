package handlers

import (
	"math"
	"net/http"
	"strconv"

	"books/pkg/models"
	"books/pkg/services"
	"books/pkg/utils"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service services.BookService
}

func NewBookHandler(service services.BookService) *BookHandler {
	return &BookHandler{service}
}

// Input Struct for Validation
type CreateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	AuthorID    uint   `json:"author_id" binding:"required"`
	PublisherID uint   `json:"publisher_id" binding:"required"`
}

func (h *BookHandler) Create(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:       input.Title,
		Description: input.Description,
		AuthorID:    input.AuthorID,
		PublisherID: input.PublisherID,
	}

	if err := h.service.CreateBook(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
}

func (h *BookHandler) List(c *gin.Context) {
	// 1. Generate Pagination from Query Params
	pagination := utils.GeneratePaginationFromRequest(c)

	// 2. Call Service
	books, totalRows, err := h.service.GetAllBooks(pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3. Calculate Metadata
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	c.JSON(http.StatusOK, gin.H{
		"data": books,
		"meta": gin.H{
			"total_rows":  totalRows,
			"total_pages": totalPages,
			"page":        pagination.Page,
			"limit":       pagination.Limit,
		},
	})
}

func (h *BookHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.service.GetBook(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}
