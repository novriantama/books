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

type PublisherHandler struct {
	service services.PublisherService
}

func NewPublisherHandler(service services.PublisherService) *PublisherHandler {
	return &PublisherHandler{service}
}

type CreatePublisherInput struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
}

func (h *PublisherHandler) Create(c *gin.Context) {
	var input CreatePublisherInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publisher := models.Publisher{Name: input.Name, Address: input.Address}
	if err := h.service.CreatePublisher(publisher); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create publisher"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Publisher created successfully"})
}

func (h *PublisherHandler) List(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	publishers, totalRows, err := h.service.GetAllPublishers(pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	c.JSON(http.StatusOK, gin.H{
		"data": publishers,
		"meta": gin.H{"total_rows": totalRows, "total_pages": totalPages, "page": pagination.Page, "limit": pagination.Limit},
	})
}

func (h *PublisherHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input CreatePublisherInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdatePublisher(uint(id), models.Publisher{Name: input.Name, Address: input.Address}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update publisher"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Publisher updated"})
}

func (h *PublisherHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeletePublisher(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Publisher deleted"})
}
