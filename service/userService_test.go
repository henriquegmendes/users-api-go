package service

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"henrique.mendes/users-api/dtos/request"
	"henrique.mendes/users-api/models"
	"henrique.mendes/users-api/utils"
)

func TestServiceCreateUser(t *testing.T) {
	s := NewUsersService(newMockUserDAO())

	users := s.FindByNamePaginated("", 1, 10)
	assert.Equal(t, 2, len(users.Data))

	user, err := s.Create(&request.UserCreateRequest{
		Name:           "Henrique",
		Age:            33,
		Email:          "henrique@henrique.com",
		Password:       "123456",
		RepeatPassword: "123456",
		Address:        "Rua Bla, 1234",
	})

	users = s.FindByNamePaginated("", 1, 10)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Henrique", user.Name)
	assert.Equal(t, "henrique@henrique.com", user.Email)
	assert.Equal(t, 3, len(users.Data))
}

func TestServiceCreateUserWithEmailAlreadyTaken(t *testing.T) {
	s := NewUsersService(newMockUserDAO())

	_, err := s.Create(&request.UserCreateRequest{
		Name:           "Henrique",
		Age:            33,
		Email:          "john.snow@winterfell.com",
		Password:       "123456",
		RepeatPassword: "123456",
		Address:        "Rua Bla, 1234",
	})

	users := s.FindByNamePaginated("", 1, 10)

	assert.NotNil(t, err)
	assert.Equal(t, 2, len(users.Data))
}

func TestServiceFindUserById(t *testing.T) {
	s := NewUsersService(newMockUserDAO())

	response := s.FindById(1)

	assert.NotNil(t, response)
	assert.Equal(t, "John Snow", response.Name)
	assert.Equal(t, "john.snow@winterfell.com", response.Email)
}

func TestServiceFindUserByIdNotFound(t *testing.T) {
	s := NewUsersService(newMockUserDAO())
	response := s.FindById(3)

	assert.Equal(t, 0, int(response.Id))
	assert.Empty(t, response)
}

func TestServiceFindUserByEmail(t *testing.T) {
	s := NewUsersService(newMockUserDAO())

	user, err := s.FindByEmail(request.UserAuthRequest{
		Email:    "john.snow@winterfell.com",
		Password: "1234",
	})

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, len(user.Token) > 0)
}

func TestServiceFindUserByEmailNotFound(t *testing.T) {
	s := NewUsersService(newMockUserDAO())

	response, err := s.FindByEmail(request.UserAuthRequest{
		Email:    "john@winterfell.com",
		Password: "1234",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "Wrong Credentials", err.Error())
	assert.Empty(t, response.Token)
}

func TestServiceFindUserByEmailWrongPassword(t *testing.T) {
	s := NewUsersService(newMockUserDAO())

	response, err := s.FindByEmail(request.UserAuthRequest{
		Email:    "john.snow@winterfell.com",
		Password: "312313123123",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "Wrong Credentials", err.Error())
	assert.Empty(t, response.Token)
}

func TestServiceFindUsersPaginated(t *testing.T) {
	s := NewUsersService(newMockUserDAO())

	users := s.FindByNamePaginated("", 1, 1)

	assert.NotNil(t, users)
	assert.Equal(t, 1, len(users.Data))
	assert.Equal(t, 2, users.Page.TotalResults)
}

func TestServiceFindUsersPaginatedWithFilter(t *testing.T) {
	s := NewUsersService(newMockUserDAO())

	users := s.FindByNamePaginated("Ne", 1, 10)

	assert.NotNil(t, users)
	assert.Equal(t, 1, len(users.Data))
	assert.Equal(t, 1, users.Page.TotalResults)
}

func TestServiceUpdateUser(t *testing.T) {
	s := NewUsersService(newMockUserDAO())

	userBeforeUpdate := s.FindById(1)
	assert.Equal(t, "John Snow", userBeforeUpdate.Name)

	userAfterUpdate := s.UpdateUserById(1, request.UserUpdateRequest{
		Name:    "Henrique",
		Age:     55,
		Address: "Summerfell",
	})

	assert.NotNil(t, "Henrique", userAfterUpdate.Name)
	assert.Equal(t, 55, userAfterUpdate.Age)
	assert.Equal(t, "Summerfell", userAfterUpdate.Address)
}

func TestServiceUpdateUserNotUpdatingEmptyFields(t *testing.T) {
	s := NewUsersService(newMockUserDAO())

	userBeforeUpdate := s.FindById(1)
	assert.Equal(t, "John Snow", userBeforeUpdate.Name)

	userAfterUpdate := s.UpdateUserById(1, request.UserUpdateRequest{
		Name:    "",
		Age:     55,
		Address: "Summerfell",
	})

	assert.NotNil(t, "John Snow", userAfterUpdate.Name)
	assert.Equal(t, 55, userAfterUpdate.Age)
	assert.Equal(t, "Summerfell", userAfterUpdate.Address)
}

func TestServiceDeleteUser(t *testing.T) {
	s := NewUsersService(newMockUserDAO())

	usersBeforeDelete := s.FindByNamePaginated("", 1, 10)
	assert.Equal(t, 2, len(usersBeforeDelete.Data))
	assert.Equal(t, 2, usersBeforeDelete.Page.TotalResults)

	s.DeleteUserById(1)

	usersAfterDelete := s.FindByNamePaginated("", 1, 10)
	assert.Equal(t, 1, len(usersAfterDelete.Data))
	assert.Equal(t, 1, usersAfterDelete.Page.TotalResults)
}

type mockUserDAO struct {
	records []models.User
	record  models.User
}

func newMockUserDAO() UsersRepository {
	encPassword, _ := utils.GenerateEncryptedPassword("1234")

	return &mockUserDAO{
		records: []models.User{
			{Id: 1, Name: "John Snow", Email: "john.snow@winterfell.com", Address: "Winterfell", Password: encPassword},
			{Id: 2, Name: "Ned Stark", Email: "ned.stark@winterfell.com", Address: "Winterfell", Password: encPassword},
		},
	}
}

func (mock *mockUserDAO) Create(user models.User) (models.User, error) {
	for i := 0; i < len(mock.records); i++ {
		if mock.records[i].Email == user.Email {
			return models.User{}, errors.New("User already exists")
		}
	}

	mock.records = append(mock.records, user)

	return user, nil
}

func (mock *mockUserDAO) FindByNamePaginated(name string, page int, limit int) ([]models.User, int) {
	var usersFiltered []models.User
	offset := (page - 1) * limit

	if len(name) > 0 {
		for i := 0; i < len(mock.records); i++ {
			fmt.Println(mock.records[i].Name, name)
			if strings.Contains(mock.records[i].Name, name) {
				usersFiltered = append(usersFiltered, mock.records[i])
			}
		}
	} else {
		usersFiltered = mock.records
	}

	if (offset + limit) >= len(usersFiltered) {
		limit = len(usersFiltered)
	} else {
		limit = offset + limit
	}

	return usersFiltered[offset:limit], len(usersFiltered)
}

func (mock *mockUserDAO) FindByEmail(email string) models.User {
	for i := 0; i < len(mock.records); i++ {
		if mock.records[i].Email == email {
			return mock.records[i]
		}
	}

	return models.User{}
}

func (mock *mockUserDAO) FindById(userId uint) models.User {
	for i := 0; i < len(mock.records); i++ {
		if mock.records[i].Id == userId {
			return mock.records[i]
		}
	}

	return models.User{}
}

func (mock *mockUserDAO) Paginate(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func (mock *mockUserDAO) UpdateUserById(userId uint64, data request.UserUpdateRequest) models.User {
	user := models.User{}
	for i := 0; i < len(mock.records); i++ {
		if mock.records[i].Id == uint(userId) {
			user = mock.records[i]
			break
		}
	}

	if user.Id == 0 {
		return models.User{}
	}

	if len(data.Name) > 0 {
		user.Name = data.Name
	}
	if data.Age > 0 {
		user.Age = data.Age
	}
	if len(data.Address) > 0 {
		user.Address = data.Address
	}

	return user
}

func (mock *mockUserDAO) DeleteUserById(userId uint64) error {
	for i := 0; i < len(mock.records); i++ {
		if mock.records[i].Id == uint(userId) {
			mock.records = append(mock.records[:i], mock.records[i+1:]...)
			return nil
		}
	}

	return errors.New("User not found")
}
