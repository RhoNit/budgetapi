package handlers

import (
	"errors"

	"github.com/RhoNit/budgetapi/internal/mailer"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	DB     *gorm.DB
	Logger echo.Logger
	Mailer mailer.Mailer
}

func (h *Handler) BindRequestBody(c echo.Context, request interface{}) error {
	if err := (&echo.DefaultBinder{}).BindBody(c, request); err != nil {
		h.Logger.Error(err)
		return errors.New("failed to bind request body... make sure you are sending a valid payload")
	}

	return nil
}
