package repository

import (
	"database/sql"
	"time"
)

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{CustomerID: 1001, Name: "Ashish", City: "New Delhi", ZipCode: "11101", DateOfBirth: time.Now()},
		{CustomerID: 1002, Name: "Ashish", City: "New Delhi", ZipCode: "11101", DateOfBirth: time.Now()},
		{CustomerID: 1003, Name: "Ashish", City: "New Delhi", ZipCode: "11101", DateOfBirth: time.Now()},
		{CustomerID: 1004, Name: "Ashish", City: "New Delhi", ZipCode: "11101", DateOfBirth: time.Now()},
	}
	return customerRepositoryMock{customers: customers}
}

func (repo customerRepositoryMock) GetAll() ([]Customer, error) {
	return repo.customers, nil
}

func (repo customerRepositoryMock) GetByID(id int) (*Customer, error) {
	for _, cust := range repo.customers {
		if cust.CustomerID == id {
			return &cust, nil
		}
	}
	return nil, sql.ErrNoRows
}
