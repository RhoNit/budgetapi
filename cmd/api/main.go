package main

import (
	"fmt"
	"os"

	"github.com/RhoNit/budgetapi/cmd/api/handlers"
	"github.com/RhoNit/budgetapi/cmd/api/middlewares"
	"github.com/RhoNit/budgetapi/cmd/api/routes"
	"github.com/RhoNit/budgetapi/common"
	"github.com/RhoNit/budgetapi/internal/mailer"
	"github.com/RhoNit/budgetapi/internal/migration"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Application struct {
	logger        echo.Logger
	server        *echo.Echo
	handler       handlers.Handler
	appMiddleware middlewares.AppMiddleware
}

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("error loading .env configs")
	}

	db, err := common.NewPostgres()
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	migration.MigrateDBUp(db)

	appMailer := mailer.NewMailer(e.Logger)

	h := handlers.Handler{
		DB:     db,
		Logger: e.Logger,
		Mailer: appMailer,
	}

	appMiddleware := middlewares.AppMiddleware{
		Logger: e.Logger,
		DB:     db,
	}

	app := Application{
		logger:        e.Logger,
		server:        e,
		handler:       h,
		appMiddleware: appMiddleware,
	}

	app.server.Use(middleware.Logger())
	// app.server.Use(middlewares.CustomMiddleware)

	routes.Endpoints(app.server, app.handler, app.appMiddleware)

	// serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	// addr := fmt.Sprintf("%s:%s", serverHost, serverPort)
	addr := fmt.Sprintf(":%s", serverPort)
	e.Logger.Fatal(e.Start(addr))
}
