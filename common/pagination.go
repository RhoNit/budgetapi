package common

import (
	"math"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Pagination struct {
	PageSize   int         `query:"page_size" json:"page_size"`
	Page       int         `query:"page" json:"page"`
	Sort       string      `query:"sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Items      interface{} `json:"items"`
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetPageSize() int {
	if p.PageSize > 100 {
		p.PageSize = 100
	} else if p.PageSize <= 0 {
		p.PageSize = 10
	}
	return p.PageSize
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPageSize()
}

func NewPaginator(model interface{}, r *http.Request, db *gorm.DB) *Pagination {
	var pagination Pagination

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("page_size"))

	var totalRows int64
	db.Model(model).Count(&totalRows)

	pagination.PageSize = pageSize
	pagination.Page = page
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetPageSize())))
	pagination.TotalPages = totalPages

	return &pagination
}

func (p *Pagination) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetPageSize())
	}
}
