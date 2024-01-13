package services

import (
	"banking/repository"
	"database/sql"
	"errors"
	"log"
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
		log.Println(err)
		return nil, err
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
			return nil, errors.New("customer not found")
		}
		return nil, err
	}
	return &CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}, nil
}
