package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/RhoNit/budgetapi/cmd/api/requests"
	"github.com/RhoNit/budgetapi/common"
	"github.com/RhoNit/budgetapi/common/custom_errors"
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

	if fetchedBudget != nil {
		return nil, errors.New(fmt.Sprintf("budget with user_id %d, title %s, year %d and month %d is already there in the database", fetchedBudget.UserID, fetchedBudget.Slug, fetchedBudget.Year, fetchedBudget.Month))
	}

	return fetchedBudget, nil
}

func (b *BudgetService) budgetExistsForYearMonthSlugAndUserID(userId uint, slug string, year uint16, month uint8) (*models.Budget, error) {
	var retrievedBudget *models.Budget

	result := b.Database.Model(&models.Budget{}).Where("user_id = ? AND year = ? AND month = ? AND slug = ?", userId, year, month, slug).First(&retrievedBudget)
	if result.Error != nil {
		return nil, result.Error
	}

	return retrievedBudget, nil
}

func (b *BudgetService) ListBudgets(query *gorm.DB, paginator *common.Pagination, budgets []*models.Budget) (*common.Pagination, error) {
	query.Scopes(paginator.Paginate()).Find(&budgets)
	paginator.Items = budgets

	return paginator, nil
}

func (b *BudgetService) GetBudgetById(budgetId uint) (*models.Budget, error) {
	budget := new(*models.Budget)

	result := b.Database.First(budget, budgetId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, custom_errors.NewCustomError("Budget not found")
		}
		return nil, errors.New("failed to fetch budget")
	}

	return *budget, nil
}

func (b *BudgetService) UpdateBudget(budget *models.Budget, budgetUpdateRequest *requests.UpdateBudgetRequest) (*models.Budget, error) {
	if budgetUpdateRequest.Date != "" {
		timeParsed, err := time.Parse(time.DateOnly, budgetUpdateRequest.Date)
		if err != nil {
			return nil, err
		}

		budget.Date = timeParsed
	}

	if budgetUpdateRequest.Amount > 0 {
		budget.Amount = budgetUpdateRequest.Amount
	}

	if budgetUpdateRequest.Description != nil {
		budget.Description = budgetUpdateRequest.Description
	}

	if budgetUpdateRequest.Title != "" {
		budget.Title = budgetUpdateRequest.Title
		slug := strings.ToLower(budgetUpdateRequest.Title)
		slug = strings.Replace(slug, " ", "_", -1)
		budget.Slug = slug
	}

	count := b.countForYearAndMonthAndSlugAndUserIDExcludingBudgetID(
		budget.UserID,
		budget.Slug,
		budget.Year,
		budget.Month,
		budget.ID,
	)
	if count > 0 {
		return nil, errors.New("budget with selected moth, year and title already exits")
	}

	b.Database.Model(&budget).Updates(budget)
	return budget, nil
}

func (b *BudgetService) countForYearAndMonthAndSlugAndUserIDExcludingBudgetID(userId uint, slug string, year uint16, month uint8, budgetId uint) int64 {
	var count int64
	b.Database.Where("user_id = ? AND slug = ? AND year = ? AND slug = ? AND id <> ?", userId, slug, year, month, budgetId).Count(&count)

	return count
}
