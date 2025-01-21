package services

import (
	"errors"
	"strings"

	"github.com/RhoNit/budgetapi/cmd/api/requests"
	"github.com/RhoNit/budgetapi/internal/models"
	"gorm.io/gorm"
)

type CategoryService struct {
	Database *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{
		Database: db,
	}
}

func (c *CategoryService) ListCategories() ([]*models.Category, error) {
	var categories []*models.Category

	result := c.Database.Find(&categories)
	if result.Error != nil {
		return nil, errors.New("failed to fetch categories")
	}

	return categories, nil
}

func (c *CategoryService) CreateCategory(categoryRequest *requests.CategoryRequest) (*models.Category, error) {
	slug := strings.ToLower(categoryRequest.Name)
	slug = strings.Replace(slug, "_", " ", -1)
	categoryCreated := &models.Category{
		Slug: slug,
		Name: categoryRequest.Name,
	}

	result := c.Database.Where(models.Category{Slug: slug}).FirstOrCreate(categoryCreated)
	if result.Error != nil {
		return nil, errors.New("failed to create category")
	}

	return categoryCreated, nil
}
