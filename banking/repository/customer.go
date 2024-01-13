package repository

import "time"

type Customer struct {
	CustomerID  int       `db:"customer_id"`
	Name        string    `db:"name"`
	DateOfBirth time.Time `db:"date_of_birth"`
	City        string    `db:"city"`
	ZipCode     string    `db:"zipcode"`
	Status      bool      `db:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetByID(int) (*Customer, error)
}
