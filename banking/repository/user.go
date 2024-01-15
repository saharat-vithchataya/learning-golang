package repository

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type UserRepository interface {
	Create(User) (*User, error)
	GetByUsername(string) (*User, error)
}
