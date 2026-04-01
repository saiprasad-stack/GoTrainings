
```md
# Go Database

## Project Structure
go-db-app/
│
├── cmd/
│   └── main.go              # Entry point
│
├── internal/
│   ├── db/
│   │   └── db.go           # DB connection setup
│   │
│   ├── models/
│   │   └── user.go         # Structs
│   │
│   ├── repository/
│   │   └── user_repo.go    # DB queries (raw SQL)
│   │
│   └── service/
│       └── user_service.go # Business logic
│
├── go.mod
```


## How It Works
### Flow
1. `main.go` initializes DB and service
2. Service layer handles business logic
3. Repository layer executes SQL queries
4. `database/sql` sends query to DB
5. DB processes and returns results
6. Results are mapped to Go structs

---

## Setup

### 1. Install PostgreSQL
Create a database:
```
CREATE DATABASE testdb;
```

### 2. Initialize project
go mod init go-db-app
go mod tidy

### 3. Run
```
go run ./cmd
```

## Code

### `cmd/main.go`

```go
package main

import (
	"fmt"
	"log"

	"go-db-app/internal/db"
	"go-db-app/internal/service"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	fmt.Println("DB Connected")

	userService := service.NewUserService(database)

	userService.CreateTable()

	userService.CreateUser("John", "john@example.com")
	userService.CreateUser("Alice", "alice@example.com")

	user := userService.GetUser(1)
	fmt.Println("User:", user)

	users := userService.GetAllUsers()
	fmt.Println("Users:", users)

	userService.UpdateUser(1, "John Updated")
	userService.DeleteUser(2)
}
```

#### `internal/db/db.go`
```go
package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=testdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
```

#### `internal/models/user.go`

```go
package models

type User struct {
	ID    int
	Name  string
	Email string
}
```

---

#### `internal/repository/user_repo.go`

```go
package repository

import (
	"database/sql"
	"go-db-app/internal/models"
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
```

---

#### `internal/service/user_service.go`

```go
package service

import (
	"database/sql"
	"log"

	"go-db-app/internal/models"
	"go-db-app/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		repo: repository.NewUserRepository(db),
	}
}

func (s *UserService) CreateTable() {
	if err := s.repo.CreateTable(); err != nil {
		log.Println(err)
	}
}

func (s *UserService) CreateUser(name, email string) {
	if err := s.repo.CreateUser(name, email); err != nil {
		log.Println(err)
	}
}

func (s *UserService) GetUser(id int) models.User {
	user, err := s.repo.GetUser(id)
	if err != nil {
		log.Println(err)
	}
	return user
}

func (s *UserService) GetAllUsers() []models.User {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		log.Println(err)
	}
	return users
}

func (s *UserService) UpdateUser(id int, name string) {
	if err := s.repo.UpdateUser(id, name); err != nil {
		log.Println(err)
	}
}

func (s *UserService) DeleteUser(id int) {
	if err := s.repo.DeleteUser(id); err != nil {
		log.Println(err)
	}
}
```
---

## Flow
```md
* Uses `database/sql` (raw SQL approach)
* Queries written inside repository layer
* Clean separation:
  * `main` → entry point
  * `service` → business logic
  * `repository` → DB logic
  * `models` → data structures
```