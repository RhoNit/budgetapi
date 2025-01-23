package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/RhoNit/budgetapi/cmd/api/requests"
	"github.com/RhoNit/budgetapi/common"
	"github.com/RhoNit/budgetapi/common/custom_errors"
	"github.com/RhoNit/budgetapi/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryService struct {
	Database *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{
		Database: db,
	}
}

func (c *CategoryService) ListCategories(paginator *common.Pagination, categories []*models.Category) (*common.Pagination, error) {
	c.Database.Scopes(paginator.Paginate()).Find(&categories)
	paginator.Items = categories
	// if result.Error != nil {
	// 	return nil, errors.New("failed to fetch categories")
	// }

	return paginator, nil
}

func (c *CategoryService) CreateCategory(categoryRequest *requests.CategoryRequest) (*models.Category, error) {
	slug := strings.ToLower(categoryRequest.Name)
	slug = strings.Replace(slug, " ", "_", -1)
	categoryCreated := &models.Category{
		Slug:     slug,
		Name:     categoryRequest.Name,
		IsCustom: categoryRequest.IsCustom,
	}

	result := c.Database.Where(models.Category{Slug: slug, Name: categoryRequest.Name}).FirstOrCreate(categoryCreated)
	if result.Error != nil {
		fmt.Println("failed to create category", slug)
		fmt.Println(result.Error.Error())
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return categoryCreated, nil
		}
		return nil, errors.New("failed to create category")
	}

	return categoryCreated, nil
}

func (c *CategoryService) GetCategoryById(idParam uint) (*models.Category, error) {
	var category *models.Category

	result := c.Database.First(&category, idParam)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, custom_errors.NewCustomError("category not found")
		}

		return nil, errors.New("failed to fetch catagory")
	}

	return category, nil
}

func (c *CategoryService) DeleteCategoryById(id uint) error {
	var category *models.Category

	category, err := c.GetCategoryById(id)
	if err != nil {
		return err
	}

	c.Database.Delete(category)

	return nil
}

func (c *CategoryService) AssociateUserToCategories(user *models.User, categories []*models.Category) error {
	if user != nil && categories != nil && len(categories) > 0 {
		var userCategories []*models.UserCategory

		for _, category := range categories {
			userCategories = append(userCategories, &models.UserCategory{
				UserID:     user.ID,
				CategoryID: category.ID,
			})
		}

		result := c.Database.Clauses(clause.OnConflict{DoNothing: true}).Create(userCategories)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (c *CategoryService) GetMultipleCategories(categoryIDs []uint) ([]*models.Category, error) {
	var categories []*models.Category
	result := c.Database.Where("id IN ?", categoryIDs).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}
