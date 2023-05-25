package service

import "fiber-test/internal/models"

//go:generate mockery --name=Interface
type Interface interface {
	GetUsers() []models.User
	GetUserByID(id int) (*models.User, error)
	CreateUser(name, email string) error
	DeleteUser(id int) error
}
