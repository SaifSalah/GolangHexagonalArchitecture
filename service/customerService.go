package service

import "saifsalah/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.ICustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {

	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepositorySub) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
