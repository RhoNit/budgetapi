package handlers

import (
	"errors"

	"github.com/RhoNit/budgetapi/cmd/api/requests"
	"github.com/RhoNit/budgetapi/cmd/api/services"
	"github.com/RhoNit/budgetapi/common"
	"github.com/RhoNit/budgetapi/common/custom_errors"
	"github.com/RhoNit/budgetapi/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateCategoryHandler(c echo.Context) error {
	_, ok := c.Get("user").(models.User)
	if !ok {
		return common.SendInternalServerErrorResponse(c, "User authentication failed")
	}

	// bind the payload body
	request := new(requests.CategoryRequest)
	if err := h.BindRequestBody(c, request); err != nil {
		return common.SendBadRequestResponse(c, "failed to bind category request body")
	}

	// validation
	validationErrors := h.ValidateRequestBody(c, *request)
	if validationErrors != nil {
		return common.SendFailedValidationResponse(c, validationErrors)
	}

	//
	categoryServices := services.NewCategoryService(h.DB)
	category, err := categoryServices.CreateCategory(request)
	if err != nil {
		return common.SendInternalServerErrorResponse(c, err.Error())
	}

	return common.SendSuccessResponse(c, "category created", category)
}

func (h *Handler) ListCategoriesHandler(c echo.Context) error {
	var categories []*models.Category
	paginator := common.NewPaginator(categories, c.Request(), h.DB)

	categoryService := services.NewCategoryService(h.DB)
	paginatedCategories, err := categoryService.ListCategories(paginator, categories)

	if err != nil {
		return common.SendInternalServerErrorResponse(c, err.Error())
	}

	return common.SendSuccessResponse(c, "categories retrieved", paginatedCategories)
}

func (h *Handler) DeleteCategoryHandler(c echo.Context) error {
	_, ok := c.Get("user").(models.User)
	if !ok {
		return common.SendInternalServerErrorResponse(c, "User authentication failed")
	}

	// bind path request
	idParamRequest := new(requests.IDParamRequest)
	if err := (&echo.DefaultBinder{}).BindPathParams(c, idParamRequest); err != nil {
		return common.SendBadRequestResponse(c, "failed to bind id path param")
	}

	// fetch category by id
	categoryService := services.NewCategoryService(h.DB)

	// perform delete operation
	err := categoryService.DeleteCategoryById(idParamRequest.ID)
	if err != nil {
		if errors.Is(err, custom_errors.NewCustomError(err.Error())) {
			return common.SendNotFoundResponse(c, err.Error())
		}
		return common.SendBadRequestResponse(c, err.Error())
	}

	return common.SendSuccessResponse(c, "category deleted", nil)
}
