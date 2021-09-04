package service

import (
	"errors"

	"gorm.io/gorm"
	"henrique.mendes/users-api/dtos/request"
	"henrique.mendes/users-api/dtos/response"
	"henrique.mendes/users-api/mappers"
	"henrique.mendes/users-api/models"
)

type UsersRepository interface {
	Paginate(page int, limit int) func(db *gorm.DB) *gorm.DB
	Create(user models.User) (models.User, error)
	FindByEmail(email string) models.User
	FindById(userId uint) models.User
	FindByNamePaginated(name string, page int, limit int) ([]models.User, int)
	UpdateUserById(userId uint64, data request.UserUpdateRequest) models.User
	DeleteUserById(userId uint64) error
}

type UsersService struct {
	repository UsersRepository
}

func NewUsersService(repository UsersRepository) *UsersService {
	return &UsersService{
		repository: repository,
	}
}

func (s *UsersService) Create(data *request.UserCreateRequest) (response.UserResponse, error) {
	if data.Password != data.RepeatPassword {
		return response.UserResponse{}, errors.New("Passwords does not match")
	}

	user, err := s.repository.Create(mappers.ToCreateUserEntity(*data))

	if err != nil {
		return response.UserResponse{}, errors.New(err.Error())
	}

	return mappers.ToUserResponse(user), nil
}

func (s *UsersService) FindByEmail(data request.UserAuthRequest) (response.UserAuthResponse, error) {
	user := s.repository.FindByEmail(data.Email)

	if user.Id == 0 || user.HasInvalidPassword(data.Password) {
		return response.UserAuthResponse{}, errors.New("Wrong Credentials")
	}

	return mappers.ToUserAuthResponse(user)
}

func (s *UsersService) FindById(userId uint) response.UserResponse {
	user := s.repository.FindById(userId)

	return mappers.ToUserResponse(user)
}

func (s *UsersService) FindByNamePaginated(name string, page int, limit int) response.UsersListResponse {
	users, total := s.repository.FindByNamePaginated(name, page, limit)

	return mappers.ToUsersListResponse(users, total, page, limit)
}

func (s *UsersService) UpdateUserById(userId uint64, data request.UserUpdateRequest) response.UserResponse {
	user := s.repository.UpdateUserById(userId, data)

	return mappers.ToUserResponse(user)
}

func (s *UsersService) DeleteUserById(userId uint64) error {
	error := s.repository.DeleteUserById(userId)

	return error
}
