package handlers

import (
	"github.com/RhoNit/budgetapi/internal/mailer"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	DB     *gorm.DB
	Logger echo.Logger
	Mailer mailer.Mailer
}
