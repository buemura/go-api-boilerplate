package routes

import (
	"github.com/buemura/go-api-boilerplate/internal/application"
	"github.com/buemura/go-api-boilerplate/internal/infra/controllers"
	"github.com/labstack/echo/v4"
)

func setupCustomerRoutes(e *echo.Echo) {
	cs := application.NewCustomerService()
	cc := controllers.NewCustomerController(cs)

	e.POST("/customer", cc.Create)
}
