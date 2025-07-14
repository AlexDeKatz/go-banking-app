package domain

import "github.com/AlexDeKatz/banking/dto"

type Customer struct {
	Id          string `db:"customer_id" json:"id" xml:"id"`
	Name        string `json:"full_name" xml:"full_name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip_code" xml:"zip_code"`
	DateOfBirth string `db:"date_of_birth" json:"date_of_birth" xml:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}

type CustomerRepository interface {
	/*
		Status == 1 means active customers
		Status == 0 means inactive customers
		Status == "" means all customers
	*/
	FindAll(string) ([]Customer, error)
	FindById(string) (*Customer, error)
}

func (c *Customer) ToDTO() *dto.Customer {
	status := "active"
	if c.Status == "0" {
		status = "inactive"
	}
	return &dto.Customer{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      status,
	}
}
