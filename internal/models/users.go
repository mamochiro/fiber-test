package models

import "errors"

var (
	UserLists []User
	IDIndex   int
)

type User struct {
	ID    int
	Name  string
	Email string
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *User) AddList(name, email string) {
	IDIndex += 1
	UserLists = append(UserLists, User{
		ID:    IDIndex,
		Name:  name,
		Email: email,
	})
}

func (c *User) GetUserByID(id int) (*User, error) {
	for _, u := range UserLists {
		if u.ID == id {
			return &User{
				ID:    u.ID,
				Name:  u.Name,
				Email: u.Email,
			}, nil
		}
	}
	return nil, errors.New("user not found")
}

func (c *User) RemoveList(id int) error {
	var checkRemove bool
	for i, u := range UserLists {
		if u.ID == id {
			UserLists = append(UserLists[:i], UserLists[i+1:]...)
			checkRemove = true
		}
	}

	if !checkRemove {
		return errors.New("user id not found")
	}

	return nil
}
