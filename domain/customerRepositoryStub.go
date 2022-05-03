package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Stephan", "Tangerang", "15144", "2000-01-01", "1"},
		{"1002", "Dwi", "Tangerang", "15144", "2002-02-02", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
