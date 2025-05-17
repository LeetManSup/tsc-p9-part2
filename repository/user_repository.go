package repository

import (
	"database/sql"
)

type User struct {
	ID   int
	Name string
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", user.ID, user.Name)
	return err
}

func (r *UserRepository) Find(id int) (*User, error) {
	row := r.db.QueryRow("SELECT id, name FROM users WHERE id = $1", id)
	var user User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
