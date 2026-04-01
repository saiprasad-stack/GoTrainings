package service

import (
	"database/sql"
	"log"

	"Database/internal/models"
	"Database/internal/repository"
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
