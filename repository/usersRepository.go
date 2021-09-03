package repository

import (
	"errors"

	"henrique.mendes/users-api/database"
	"henrique.mendes/users-api/models"
)

func Create(user models.User) (models.User, error) {
	err := database.DB.Create(&user).Error

	if err != nil {
		return models.User{}, errors.New(err.Error())
	}

	return user, nil
}

func FindByEmail(email string) models.User {
	var user models.User
	database.DB.Where("email = ?", email).First(&user)

	return user
}
