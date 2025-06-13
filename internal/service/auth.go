package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/kirakazza/go-todo-api/internal/model"
	"github.com/kirakazza/go-todo-api/internal/repository"
)

func RegisterUser(username, password string) error {
	// Проверим, есть ли такой юзер
	existing, _ := repository.GetUserByUsername(username)
	if existing != nil {
		return errors.New("user already exists")
	}

	// Хешируем пароль
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Username:     username,
		PasswordHash: string(hash),
	}

	// Сохраняем
	return repository.CreateUser(user)
}

func AuthenticateUser(username, password string) (*model.User, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
