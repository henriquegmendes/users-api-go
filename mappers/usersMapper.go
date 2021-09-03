package mappers

import (
	"golang.org/x/crypto/bcrypt"
	"henrique.mendes/users-api/dtos/request"
	"henrique.mendes/users-api/dtos/response"
	"henrique.mendes/users-api/models"
	"henrique.mendes/users-api/utils"
)

func ToCreateUserEntity(data request.UserCreateRequest) models.User {
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)

	return models.User{
		Name:     data.Name,
		Age:      data.Age,
		Email:    data.Email,
		Password: encryptedPassword,
		Address:  data.Address,
	}
}

func ToCreateUserResponse(user models.User) response.UserCreateResponse {
	return response.UserCreateResponse{
		Id:      user.Id,
		Name:    user.Name,
		Age:     user.Age,
		Email:   user.Email,
		Address: user.Address,
	}
}

func ToUserAuthResponse(user models.User) (response.UserAuthResponse, error) {
	token, error := utils.GenerateUserJwt(user)

	if error != nil {
		return response.UserAuthResponse{}, error
	}

	return response.UserAuthResponse{Token: token}, nil
}
