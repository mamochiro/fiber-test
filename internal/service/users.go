package service

import "fiber-test/internal/models"

var user models.User

func (s Service) GetUsers() []models.User {
	return models.UserLists
}

func (s Service) GetUserByID(id int) (*models.User, error) {
	data, err := user.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s Service) CreateUser(name, email string) error {
	user.AddList(name, email)
	return nil
}

func (s Service) DeleteUser(id int) error {
	if err := user.RemoveList(id); err != nil {
		return err
	}
	return nil
}
