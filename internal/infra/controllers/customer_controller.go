package controllers

import (
	"net/http"

	"github.com/buemura/go-api-boilerplate/internal/application"
	"github.com/buemura/go-api-boilerplate/internal/domain/customer"
	"github.com/labstack/echo/v4"
)

type CustomerController struct {
	customerService *application.CustomerService
}

func NewCustomerController(customerService *application.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: customerService,
	}
}

func (cc *CustomerController) Create(c echo.Context) error {
	input := customer.ICreateCustomer{
		Name:  "haha",
		Email: "haha@haha.com",
	}
	cus, err := cc.customerService.Create(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, cus)
}
