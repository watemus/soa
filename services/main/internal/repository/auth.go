package repository

import (
	"main-service/internal/server"
)

type Hashed string

type AuthRepository interface {
	CreateUser(username string, password Hashed) error
	ValidateUser(username string, password Hashed) (bool, error)
	UpdateUser(username string, body *server.UpdateUserBody) error
	DoesUserExist(username string) (bool, error)
}
