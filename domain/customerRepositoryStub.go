package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindALL() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "111", Name: "John", City: "Yale", Zipcode: "101032", DateofBirth: "2020-2-2", Status: "1"},
		{Id: "112", Name: "Doe", City: "Ukraine", Zipcode: "101342", DateofBirth: "2020-6-2", Status: "0"},
	}
	return CustomerRepositoryStub{customers: customers}
}
