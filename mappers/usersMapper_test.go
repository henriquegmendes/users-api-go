package mappers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"henrique.mendes/users-api/dtos/request"
	"henrique.mendes/users-api/models"
)

func TestToCreateUserEntity(t *testing.T) {
	request := request.UserCreateRequest{
		Name:           "Henrique",
		Age:            33,
		Email:          "henrique@henrique.com",
		Password:       "12345",
		RepeatPassword: "12345",
		Address:        "Rua bla, 1234",
	}

	mapped := ToCreateUserEntity(request)

	assert.Equal(t, request.Name, mapped.Name)
	assert.Equal(t, request.Age, mapped.Age)
	assert.Equal(t, request.Email, mapped.Email)
	assert.Equal(t, request.Address, mapped.Address)
	assert.NotEqual(t, request.Password, mapped.Password)
	assert.False(t, mapped.HasInvalidPassword("12345"))
}

func TestToUserResponse(t *testing.T) {
	user := models.User{
		Id:       1,
		Name:     "Henrique",
		Age:      33,
		Email:    "henrique@henrique.com",
		Password: []byte("12345"),
		Address:  "Rua bla, 1234",
	}

	mapped := ToUserResponse(user)

	assert.Equal(t, user.Id, mapped.Id)
	assert.Equal(t, user.Name, mapped.Name)
	assert.Equal(t, user.Age, mapped.Age)
	assert.Equal(t, user.Email, mapped.Email)
	assert.Equal(t, user.Address, mapped.Address)
}

func TestToUserAuthResponse(t *testing.T) {
	user := models.User{
		Id:       1,
		Name:     "Henrique",
		Age:      33,
		Email:    "henrique@henrique.com",
		Password: []byte("12345"),
		Address:  "Rua bla, 1234",
	}

	mapped, _ := ToUserAuthResponse(user)

	assert.True(t, len(mapped.Token) > 0)
}

func TestToUsersListResponse(t *testing.T) {
	users := []models.User{
		{
			Id:       1,
			Name:     "Henrique",
			Age:      33,
			Email:    "henrique@henrique.com",
			Password: []byte("12345"),
			Address:  "Rua bla, 1234",
		},
	}
	total := 10
	page := 1
	limit := 5

	response := ToUsersListResponse(users, total, page, limit)

	assert.Equal(t, "Henrique", response.Data[0].Name)
	assert.Equal(t, page, response.Page.Page)
	assert.Equal(t, len(users), response.Page.TotalPerPage)
	assert.Equal(t, total, response.Page.TotalResults)
	assert.Equal(t, 2, response.Page.LastPage)
}

func TestToUpdateUser(t *testing.T) {
	data := request.UserUpdateRequest{
		Name:    "Henrique",
		Age:     33,
		Address: "Rua teste",
	}

	userToUpdate := ToUpdateUser(data)

	assert.Equal(t, "Henrique", userToUpdate.Name)
	assert.Equal(t, 33, userToUpdate.Age)
	assert.Equal(t, "Rua teste", userToUpdate.Address)
}
