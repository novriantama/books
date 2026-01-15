package handlers

import (
	"books/pkg/models"
	"books/pkg/services"
	"books/pkg/utils"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	service services.AuthorService
}

func NewAuthorHandler(service services.AuthorService) *AuthorHandler {
	return &AuthorHandler{service}
}

type CreateAuthorInput struct {
	Name string `json:"name" binding:"required"`
	Bio  string `json:"bio"`
}

func (h *AuthorHandler) Create(c *gin.Context) {
	var input CreateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author := models.Author{Name: input.Name, Bio: input.Bio}
	if err := h.service.CreateAuthor(author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create author"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Author created successfully"})
}

func (h *AuthorHandler) List(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	authors, totalRows, err := h.service.GetAllAuthors(pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	c.JSON(http.StatusOK, gin.H{
		"data": authors,
		"meta": gin.H{"total_rows": totalRows, "total_pages": totalPages, "page": pagination.Page, "limit": pagination.Limit},
	})
}

func (h *AuthorHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input CreateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateAuthor(uint(id), models.Author{Name: input.Name, Bio: input.Bio}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update author"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Author updated"})
}

func (h *AuthorHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteAuthor(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
}
