package service

import (
	"errors"

	"henrique.mendes/users-api/dtos/request"
	"henrique.mendes/users-api/mappers"
	"henrique.mendes/users-api/models"
	"henrique.mendes/users-api/repository"
)

func Create(data *request.UserCreateRequest) (models.User, error) {
	if data.Password != data.RepeatPassword {
		return models.User{}, errors.New("Passwords does not match")
	}

	user, err := repository.Create(mappers.ToCreateUserEntity(*data))

	if err != nil {
		return models.User{}, errors.New(err.Error())
	}

	return user, nil
}

func FindByEmail(email string) models.User {
	return repository.FindByEmail(email)
}
