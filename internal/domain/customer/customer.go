package customer

import "github.com/google/uuid"

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewCustomer(input ICreateCustomer) (*Customer, error) {
	return &Customer{
		ID:    uuid.NewString(),
		Name:  input.Name,
		Email: input.Email,
	}, nil
}
