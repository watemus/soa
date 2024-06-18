package repository

import (
	"errors"
	"main-service/internal/server"
)

type AuthInMemory struct {
	users []User
}

func NewAuthInMemory() *AuthInMemory {
	return &AuthInMemory{}
}

func (r *AuthInMemory) CreateUser(username string, password Hashed) error {
	for _, user := range r.users {
		if user.Username == username {
			return errors.New("user exists")
		}
	}
	r.users = append(r.users, User{
		Username:     username,
		PasswordHash: password,
	})
	return nil
}

func (r *AuthInMemory) ValidateUser(username string, password Hashed) (bool, error) {
	for _, user := range r.users {
		if user.Username == username && user.PasswordHash == password {
			return true, nil
		}
	}
	return false, nil
}

func (r *AuthInMemory) UpdateUser(username string, body *server.UpdateUserBody) error {
	for _, user := range r.users {
		if user.Username == username {
			user.FirstName = body.FirstName
			user.LastName = body.LastName
			user.BirthDate = body.BirthDate
			user.Email = body.Email
			user.PhoneNumber = body.PhoneNumber
			return nil
		}
	}
	return errors.New("user not found")
}
