package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
	"main-service/internal/server"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(username string, password Hashed) error {
	query := "INSERT INTO users (username, password_hash) VALUES ($1, $2)"
	_, err := r.db.Exec(query, username, password)
	if err != nil {
		log.Printf("error while inserting user %s", err.Error())
		return errors.New("error while creating user")
	}
	return nil
}

func (r *AuthPostgres) ValidateUser(username string, password Hashed) (bool, error) {
	var user User
	query := "SELECT id FROM users WHERE username=$1 AND password_hash=$2"
	err := r.db.Get(&user, query, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("user not found")
		}
		log.Printf("error while validating user %s", err.Error())
		return false, errors.New("error while validating user")
	}
	return true, nil
}

func (r *AuthPostgres) UpdateUser(username string, body *server.UpdateUserBody) error {
	query := "UPDATE users SET first_name=$2,  last_name=$3, birth_date=$4, email=$5, phone_number=$6 WHERE username=$1"
	_, err := r.db.Exec(query, username, body.FirstName, body.LastName, body.BirthDate, body.Email, body.PhoneNumber)
	if err != nil {
		log.Printf("error while updating user %s", err.Error())
		return errors.New("error while updating user")
	}
	return nil
}
