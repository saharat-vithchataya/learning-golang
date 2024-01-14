package services

import (
	"banking/logs"
	"banking/repository"
	"database/sql"
	"errors"
)

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrUnexpectedError  = errors.New("unexpected error")
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

func (srv customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := srv.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, ErrUnexpectedError
	}
	response := []CustomerResponse{}
	for _, cust := range customers {
		custResponse := CustomerResponse{
			CustomerID: cust.CustomerID,
			Name:       cust.Name,
			Status:     cust.Status,
		}
		response = append(response, custResponse)
	}
	return response, nil
}

func (srv customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := srv.custRepo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Error(err)
			return nil, ErrCustomerNotFound
		}
		logs.Error(err)
		return nil, ErrUnexpectedError
	}
	return &CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}, nil
}
