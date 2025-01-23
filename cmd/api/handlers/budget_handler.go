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
		return common.SendInternalServerErrorResponse(c, "Budget could not be created")
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
