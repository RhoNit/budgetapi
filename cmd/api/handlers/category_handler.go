package handlers

import (
	"github.com/RhoNit/budgetapi/cmd/api/services"
	"github.com/RhoNit/budgetapi/common"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateCategoryHandler(c echo.Context) error {

	return nil
}

func (h *Handler) ListCategoriesHandler(c echo.Context) error {
	categoryService := services.NewCategoryService(h.DB)
	retrievedCategories, err := categoryService.ListCategories()

	if err != nil {
		return common.SendInternalServerErrorResponse(c, err.Error())
	}

	return common.SendSuccessResponse(c, "categories retrieved", retrievedCategories)
}
