package main

import (
	"log/slog"
	"os"

	"github.com/buemura/go-api-boilerplate/internal/infra/routes"
	"github.com/labstack/echo/v4"
)

func setupServerMiddlewares(e *echo.Echo) {
	routes.SetupRoutes(e)
}

func main() {
	// Logger
	slogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(slogger)

	// Http Server
	e := echo.New()
	setupServerMiddlewares(e)
	err := e.Start(":8080")
	if err != nil {
		slog.Error(err.Error())
	}
}
