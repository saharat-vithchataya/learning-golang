package repository

import "github.com/jmoiron/sqlx"

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return userRepository{db: db}
}

func (repo userRepository) Create(user User) (*User, error) {
	query := "INSERT INTO userss(username, password) values($1,$2) RETURNING id"
	row := repo.db.QueryRow(query, user.Username, user.Password)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo userRepository) GetByUsername(username string) (*User, error) {
	query := "select id, username, password from userss where username=$1"
	user := User{}
	err := repo.db.Get(&user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
