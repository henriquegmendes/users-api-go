package service

import (
	"errors"

	"henrique.mendes/users-api/dtos/request"
	"henrique.mendes/users-api/dtos/response"
	"henrique.mendes/users-api/mappers"
	"henrique.mendes/users-api/repository"
)

func Create(data *request.UserCreateRequest) (response.UserResponse, error) {
	if data.Password != data.RepeatPassword {
		return response.UserResponse{}, errors.New("Passwords does not match")
	}

	user, err := repository.Create(mappers.ToCreateUserEntity(*data))

	if err != nil {
		return response.UserResponse{}, errors.New(err.Error())
	}

	return mappers.ToUserResponse(user), nil
}

func FindByEmail(data request.UserAuthRequest) (response.UserAuthResponse, error) {
	user := repository.FindByEmail(data.Email)

	if user.Id == 0 || user.HasInvalidPassword(data.Password) {
		return response.UserAuthResponse{}, errors.New("Wrong credentials")
	}

	return mappers.ToUserAuthResponse(user)
}

func FindById(userId uint) response.UserResponse {
	user := repository.FindById(userId)

	return mappers.ToUserResponse(user)
}

func FindByNamePaginated(name string, page int, limit int) response.UsersListResponse {
	users, total := repository.FindByNamePaginated(name, page, limit)

	return mappers.ToUsersListResponse(users, total, page, limit)
}

func UpdateUserById(userId uint64, data request.UserUpdateRequest) response.UserResponse {
	user := repository.UpdateUserById(userId, data)

	return mappers.ToUserResponse(user)
}

func DeleteUserById(userId uint64) error {
	error := repository.DeleteUserById(userId)

	return error
}
