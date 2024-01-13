package repository

import "github.com/jmoiron/sqlx"

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepository {
	return customerRepositoryDB{db: db}
}

func (repo customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	err := repo.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (repo customerRepositoryDB) GetByID(customerID int) (*Customer, error) {
	var customer Customer
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id=$1"
	err := repo.db.Get(&customer, query, customerID)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
