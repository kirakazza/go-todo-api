package repository

import (
	"github.com/kirakazza/go-todo-api/internal/model"
)

func CreateUser(user *model.User) error {
	return DB.Create(user).Error
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	result := DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
