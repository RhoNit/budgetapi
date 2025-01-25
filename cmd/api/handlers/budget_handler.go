package handlers

import (
	"github.com/RhoNit/budgetapi/cmd/api/requests"
	"github.com/RhoNit/budgetapi/cmd/api/services"
	"github.com/RhoNit/budgetapi/common"
	"github.com/RhoNit/budgetapi/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateBudgetHandler(c echo.Context) error {
	user, ok := c.Get("user").(models.User)
	if !ok {
		return common.SendInternalServerErrorResponse(c, "User authentication failed")
	}

	// bind the payload body
	request := new(requests.CreateBudgetRequest)
	if err := h.BindRequestBody(c, request); err != nil {
		return common.SendBadRequestResponse(c, "failed to bind budget request body")
	}

	// validation
	validationErrors := h.ValidateRequestBody(c, *request)
	if validationErrors != nil {
		return common.SendFailedValidationResponse(c, validationErrors)
	}

	budgetService := services.NewBudgetService(h.DB)
	budgetCreated, err := budgetService.CreateBudget(request, user.ID)
	if err != nil {
		c.Logger().Error(err)
		return common.SendInternalServerErrorResponse(c, err.Error())
	}

	categoryService := services.NewCategoryService(h.DB)
	listOfCategories, err := categoryService.GetMultipleCategories(request.CategoryIDs)
	if err != nil {
		return common.SendInternalServerErrorResponse(c, "Budget could not be created")
	}

	// associate listOfCategories instance to budgetCreated model instance
	err = budgetService.Database.Model(budgetCreated).Association("Categories").Replace(listOfCategories)
	if err != nil {
		return common.SendInternalServerErrorResponse(c, "Budget could not be created")
	}

	budgetCreated.Categories = listOfCategories

	return common.SendSuccessResponse(c, "budget created", budgetCreated)
}

func (h *Handler) ListBudgetsHandler(c echo.Context) error {
	user, ok := c.Get("user").(models.User)
	if !ok {
		return common.SendInternalServerErrorResponse(c, "User authentication failed")
	}

	var budgets []*models.Budget

	query := h.DB.Preload("Categories").Scopes(common.WhereUserIDScope(user.ID))
	paginator := common.NewPaginator(budgets, c.Request(), query)

	budgetService := services.NewBudgetService(h.DB)
	paginatedBudgets, err := budgetService.ListBudgets(query, paginator, budgets)

	if err != nil {
		return common.SendInternalServerErrorResponse(c, err.Error())
	}

	return common.SendSuccessResponse(c, "budgets retrieved", paginatedBudgets)
}

// func (h *Handler) UpdateBudgetHandler(c echo.Context) error {
// 	user, ok := c.Get("user").(models.User)
// 	if !ok {
// 		return common.SendInternalServerErrorResponse(c, "User authentication failed")
// 	}

// 	// create request type as IDParamRequest
// 	// and bind the request body
// 	budgetID := new(requests.IDParamRequest)
// 	if err := (&echo.DefaultBinder{}).BindPathParams(c, budgetID.ID); err != nil {
// 		return common.SendBadRequestResponse(c, "failed to bind id path param")
// 	}

// 	// retrieve the budget by id
// 	budgetService := services.NewBudgetService(h.DB)
// 	retrievedBudget, err := budgetService.GetBudgetById(budgetID.ID)
// 	if err != nil {
// 		if errors.Is(err, custom_errors.NewCustomError(err.Error())) {
// 			return common.SendNotFoundResponse(c, err.Error())
// 		}
// 		return common.SendBadRequestResponse(c, err.Error())
// 	}

// 	// bind request body of type UpdateBudgetRequest
// 	request := new(requests.UpdateBudgetRequest)
// 	if err := h.BindRequestBody(c, request); err != nil {
// 		return common.SendBadRequestResponse(c, "failed to bind budget request body")
// 	}

// 	// validation
// 	validationErrors := h.ValidateRequestBody(c, request)
// 	if validationErrors != nil {
// 		return common.SendFailedValidationResponse(c, validationErrors)
// 	}

// 	budget, err := budgetService.UpdateBudget(retrievedBudget, request)
// 	if err != nil {
// 		return common.Send
// 	}

// 	return nil
// }
