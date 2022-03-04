package domain

type CustomerRepositorySub struct {
	customers []Customer
}

func (s CustomerRepositorySub) FindAll() ([]Customer, error) {

	return s.customers, nil
}

func NewCustomerRepository() CustomerRepositorySub {

	customers := []Customer{
		{Id: "01001", Name: "Saif Salah", City: "Baghdad", Zipcode: "00964", DateOfBirth: "1992-09-01", Status: ":/"},
	}

	return CustomerRepositorySub{
		customers: customers,
	}
}
