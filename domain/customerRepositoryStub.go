package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (crs *CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return crs.customers, nil
}

func NewCustomerRepositoryStub() *CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Alice", City: "Wonderland", ZipCode: "12345", DateOfBirth: "2000-01-01", Status: "active"},
		{Id: "1002", Name: "Bob", City: "Builderland", ZipCode: "67890", DateOfBirth: "1995-05-05", Status: "inactive"},
	}

	return &CustomerRepositoryStub{customers}
}
