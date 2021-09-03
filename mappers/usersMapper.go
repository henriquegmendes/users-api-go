package mappers

import (
	"math"

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

func ToUserResponse(user models.User) response.UserResponse {
	return response.UserResponse{
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

func ToUsersListResponse(users []models.User, total int, page int, limit int) response.UsersListResponse {
	var userResponse []response.UserResponse

	for i := 0; i < len(users); i++ {
		userResponse = append(userResponse, ToUserResponse(users[i]))
	}

	return response.UsersListResponse{
		Data: userResponse,
		Page: response.Page{
			Page:         page,
			TotalPerPage: len(userResponse),
			TotalResults: int(total),
			LastPage:     int(math.Ceil(float64(total) / float64(limit))),
		},
	}
}

func ToUpdateUser(data request.UserUpdateRequest) models.User {
	return models.User{
		Name:    data.Name,
		Age:     data.Age,
		Address: data.Address,
	}
}
