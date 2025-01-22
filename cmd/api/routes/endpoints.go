package routes

import (
	"net/http"

	"github.com/RhoNit/budgetapi/cmd/api/handlers"
	"github.com/RhoNit/budgetapi/cmd/api/middlewares"
	"github.com/labstack/echo/v4"
)

func Endpoints(engine *echo.Echo, handler handlers.Handler, mw middlewares.AppMiddleware) {
	appRoute := engine.Group("/api")

	publicAuthRoutes := appRoute.Group("/auth")
	{
		publicAuthRoutes.POST("/register", handler.RegisterUserHandler)
		publicAuthRoutes.POST("/login", handler.LoginUserHandler)
	}

	profileRoutes := appRoute.Group("/profile")
	{
		profileRoutes.GET("/authenticated-user", handler.GetAuthenticatedUser, mw.AuthMiddleware)
	}

	categoryRoutes := appRoute.Group("/categories", mw.AuthMiddleware)
	{
		categoryRoutes.POST("/create", handler.CreateCategoryHandler)
		categoryRoutes.GET("/all", handler.ListCategoriesHandler)
		categoryRoutes.DELETE("/delete/:id", handler.DeleteCategoryHandler)
	}

	engine.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Budget API's `Hello World` page")
	})
	engine.GET("/health", handler.HealthCheck)
}
