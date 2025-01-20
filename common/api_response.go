package common

import (
	"net/http"

	"github.com/RhoNit/budgetapi/cmd/api/validation"
	"github.com/labstack/echo/v4"
)

type ApiResponse map[string]interface{}

type JSONSuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JSONFailedValidationResponse struct {
	Success bool                          `json:"success"`
	Message string                        `json:"message"`
	Errors  []*validation.ValidationError `json:"errors"`
}

type JSONErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SendSuccessResponse(c echo.Context, msg string, data interface{}) error {
	return c.JSON(http.StatusOK, JSONSuccessResponse{
		Success: true,
		Message: msg,
		Data:    data,
	})
}

func SendFailedValidationResponse(c echo.Context, errors []*validation.ValidationError) error {
	return c.JSON(http.StatusUnprocessableEntity, JSONFailedValidationResponse{
		Success: false,
		Errors:  errors,
	})
}

func SendErrorResponse(c echo.Context, msg string, statusCode int) error {
	return c.JSON(statusCode, JSONErrorResponse{
		Success: false,
		Message: msg,
	})
}

func SendBadRequestResponse(c echo.Context, msg string) error {
	return SendErrorResponse(c, msg, http.StatusBadRequest)
}

func SendInternalServerErrorResponse(c echo.Context, msg string) error {
	return SendErrorResponse(c, msg, http.StatusInternalServerError)
}

func SendNotFoundResponse(c echo.Context, msg string) error {
	return SendErrorResponse(c, msg, http.StatusNotFound)
}
