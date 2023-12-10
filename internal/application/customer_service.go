package application

import (
	"log/slog"

	"github.com/buemura/go-api-boilerplate/internal/domain/customer"
)

type CustomerService struct {
	// Add repositories and services here
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (cs *CustomerService) Create(input customer.ICreateCustomer) (*customer.Customer, error) {
	slog.Info("[CustomerService][Create] - Creating customer")
	cus, err := customer.NewCustomer(input)
	if err != nil {
		slog.Error("[CustomerService][Create] - Error:" + err.Error())
		return nil, err
	}
	return cus, nil
}
