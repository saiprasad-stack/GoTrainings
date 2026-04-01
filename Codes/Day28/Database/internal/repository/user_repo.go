package repository

import (
	"Database/internal/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT
	);`
	_, err := r.DB.Exec(query)
	return err
}

func (r *UserRepository) CreateUser(name, email string) error {
	query := `INSERT INTO users(name, email) VALUES($1, $2)`
	_, err := r.DB.Exec(query, name, email)
	return err
}

func (r *UserRepository) GetUser(id int) (models.User, error) {
	query := `SELECT id, name, email FROM users WHERE id=$1`
	var user models.User
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(id int, name string) error {
	query := `UPDATE users SET name=$1 WHERE id=$2`
	_, err := r.DB.Exec(query, name, id)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.DB.Exec(query, id)
	return err
}
