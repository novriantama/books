package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Sort   string `json:"sort"`
	Search string `json:"search"`
}

func GeneratePaginationFromRequest(c *gin.Context) Pagination {
	limit := 10
	page := 1
	sort := "created_at desc"
	search := ""

	queryLimit, _ := strconv.Atoi(c.Query("limit"))
	queryPage, _ := strconv.Atoi(c.Query("page"))
	querySort := c.Query("sort")
	querySearch := c.Query("search")

	if queryLimit > 0 {
		limit = queryLimit
	}
	if queryPage > 0 {
		page = queryPage
	}
	if querySort != "" {
		sort = querySort
	}
	if querySearch != "" {
		search = querySearch
	}

	return Pagination{
		Limit:  limit,
		Page:   page,
		Sort:   sort,
		Search: search,
	}
}

// Scope for GORM to apply pagination and search
func (p *Pagination) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (p.Page - 1) * p.Limit
		return db.Offset(offset).Limit(p.Limit).Order(p.Sort)
	}
}
