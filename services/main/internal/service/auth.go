package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"main-service/internal/repository"
	"main-service/internal/server"
	"time"
)

const (
	salt = "ksdjflkdflskdjfl"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	tokenString, err := token.SignedString([]byte("aboba"))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJWT(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error in parsing")
		}
		return []byte("aboba"), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token error")
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		return "", errors.New("token error")
	}
	username, found := claims["username"]
	if !found {
		return "", errors.New("unknown user")
	}
	return username.(string), nil
}

func hashPassword(password string) repository.Hashed {
	hash := sha1.New()
	hash.Write([]byte(password))
	return repository.Hashed(fmt.Sprintf("%x", hash.Sum([]byte(salt))))
}

func (s *AuthService) SignIn(username string, password string) (string, error) {
	validated, err := s.repo.ValidateUser(username, hashPassword(password))
	if err != nil {
		return "", err
	}
	if !validated {
		return "", errors.New("user not exists")
	}
	token, err := GenerateJWT(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) SignUp(username string, password string) error {
	err := s.repo.CreateUser(username, hashPassword(password))
	return err
}

func (s *AuthService) UpdateUser(username string, body *server.UpdateUserBody) error {
	err := s.repo.UpdateUser(username, body)
	return err
}

func (s *AuthService) DoesUserExist(username string) (bool, error) {
	return s.repo.DoesUserExist(username)
}
