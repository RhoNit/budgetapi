package routes

import (
	"net/http"

	"github.com/RhoNit/budgetapi/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func Endpoints(engine *echo.Echo, handler handlers.Handler) {
	engine.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Budget API's `Hello World` page")
	})
	engine.GET("/health", handler.HealthCheck)
	engine.POST("/register", handler.RegisterUserHandler)
	engine.POST("/login", handler.LoginUserHandler)
}
