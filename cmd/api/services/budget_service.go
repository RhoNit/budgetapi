package services

import (
	"errors"
	"strings"
	"time"

	"github.com/RhoNit/budgetapi/cmd/api/requests"
	"github.com/RhoNit/budgetapi/internal/models"
	"gorm.io/gorm"
)

type BudgetService struct {
	Database *gorm.DB
}

func NewBudgetService(db *gorm.DB) *BudgetService {
	return &BudgetService{
		Database: db,
	}
}

func (b *BudgetService) CreateBudget(budgetRequest *requests.CreateBudgetRequest, userId uint) (*models.Budget, error) {
	slug := strings.ToLower(budgetRequest.Title)
	slug = strings.Replace(slug, " ", "_", -1)

	budget := &models.Budget{
		Amount:      budgetRequest.Amount,
		UserID:      userId,
		Title:       budgetRequest.Title,
		Slug:        slug,
		Description: budgetRequest.Description,
	}

	if budgetRequest.Date == "" {
		currentDate := time.Now()
		budget.Date = currentDate
	}

	budgetMonth := uint8(budget.Date.Month())
	budgetYear := uint16(budget.Date.Year())

	budget.Month = budgetMonth
	budget.Year = budgetYear

	fetchedBudget, err := b.budgetExistsForYearMonthSlugAndUserID(budget.UserID, budget.Slug, budget.Year, budget.Month)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := b.Database.Create(budget)
			if result.Error != nil {
				return nil, result.Error
			}
			return budget, nil
		}
		return nil, err
	}

	return fetchedBudget, nil
}

func (b *BudgetService) budgetExistsForYearMonthSlugAndUserID(userId uint, slug string, year uint16, month uint8) (*models.Budget, error) {
	var retrievedBudget *models.Budget

	result := b.Database.Where("user_id = ? AND year = ? AND month = ? AND slug = ?", userId, year, month, slug).First(&retrievedBudget)
	if result.Error != nil {
		return nil, result.Error
	}

	return retrievedBudget, nil
}
