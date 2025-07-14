package service

import (
	"github.com/AlexDeKatz/banking/domain"
	"github.com/AlexDeKatz/banking/dto"
)

type CustomerService interface {
	GetAllCustomers(string) ([]dto.Customer, error)
	GetCustomerById(string) (*dto.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (dcs DefaultCustomerService) GetAllCustomers(status string) ([]dto.Customer, error) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customer, err := dcs.repo.FindAll(status)

	if err != nil {
		return nil, err
	}
	dtoCustomers := make([]dto.Customer, 0, len(customer))
	for _, c := range customer {
		dtoCustomers = append(dtoCustomers, *c.ToDTO())
	}
	return dtoCustomers, nil
}

func (dcs DefaultCustomerService) GetCustomerById(id string) (*dto.Customer, error) {
	customer, err := dcs.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	dtoCustomer := customer.ToDTO()

	return dtoCustomer, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
