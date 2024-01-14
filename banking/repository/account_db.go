package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (repo accountRepositoryDB) Create(acc Account) (*Account, error) {
	query := "insert into accounts(customer_id, opening_date, account_type, amount, status) values ($1,$2,$3,$4,$5) RETURNING account_id"
	row := repo.db.QueryRow(query,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	)

	err := row.Scan(&acc.AccountID)
	if err != nil {
		return nil, err
	}

	return &acc, nil
}

func (repo accountRepositoryDB) GetAll(id int) ([]Account, error) {
	var accounts []Account
	query := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where customer_id = $1"
	err := repo.db.Select(&accounts, query, id)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
