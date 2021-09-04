package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAuthRequestValidationOk(t *testing.T) {
	request := UserAuthRequest{
		Email:    "henrique@henrique.com",
		Password: "123456",
	}

	validation := request.ValidateUserAuthRequest()

	assert.Empty(t, validation.Get("email"))
	assert.Empty(t, validation.Get("password"))
}

func TestUserAuthRequestValidationWithEmptyValues(t *testing.T) {
	request := UserAuthRequest{
		Email:    "",
		Password: "",
	}

	validation := request.ValidateUserAuthRequest()

	assert.NotEmpty(t, validation.Get("email"))
	assert.Equal(t, "The email field is required", validation.Get("email"))
	assert.NotEmpty(t, validation.Get("password"))
	assert.Equal(t, "The password field is required", validation.Get("password"))
}

func TestUserAuthRequestValidationWithInvalidValues(t *testing.T) {
	request := UserAuthRequest{
		Email:    "this is not an email",
		Password: "1234",
	}

	validation := request.ValidateUserAuthRequest()

	assert.NotEmpty(t, validation.Get("email"))
	assert.Equal(t, "The email field must be a valid email address", validation.Get("email"))
	assert.NotEmpty(t, validation.Get("password"))
	assert.Equal(t, "The password field must be minimum 6 char", validation.Get("password"))
}

func TestUserCreateRequestValidationOk(t *testing.T) {
	request := UserCreateRequest{
		Name:           "Henrique",
		Age:            33,
		Email:          "henrique@henrique.com",
		Password:       "123456",
		RepeatPassword: "123456",
		Address:        "Rua Bla, 1234",
	}

	validation := request.ValidateUserCreateRequest()

	assert.Empty(t, validation.Get("name"))
	assert.Empty(t, validation.Get("age"))
	assert.Empty(t, validation.Get("email"))
	assert.Empty(t, validation.Get("password"))
	assert.Empty(t, validation.Get("repeat_password"))
	assert.Empty(t, validation.Get("address"))
}

func TestUserCreateRequestValidationWithEmptyValues(t *testing.T) {
	request := UserCreateRequest{
		Name:           "",
		Email:          "",
		Password:       "",
		RepeatPassword: "",
		Address:        "",
	}

	validation := request.ValidateUserCreateRequest()

	assert.NotEmpty(t, validation.Get("name"))
	assert.Equal(t, "The name field is required", validation.Get("name"))
	assert.NotEmpty(t, validation.Get("age"))
	assert.Equal(t, "The age field is required", validation.Get("age"))
	assert.NotEmpty(t, validation.Get("email"))
	assert.Equal(t, "The email field is required", validation.Get("email"))
	assert.NotEmpty(t, validation.Get("password"))
	assert.Equal(t, "The password field is required", validation.Get("password"))
	assert.NotEmpty(t, validation.Get("repeat_password"))
	assert.Equal(t, "The repeat_password field is required", validation.Get("repeat_password"))
}

func TestUserCreateRequestValidationWithInvalidValues(t *testing.T) {
	request := UserCreateRequest{
		Name:           "He",
		Email:          "this is not an email",
		Password:       "1234",
		RepeatPassword: "1234",
		Address:        "Rua Bla, 1234",
	}

	validation := request.ValidateUserCreateRequest()

	assert.NotEmpty(t, validation.Get("name"))
	assert.Equal(t, "The name field must be minimum 3 char", validation.Get("name"))
	assert.NotEmpty(t, validation.Get("age"))
	assert.Equal(t, "The age field is required", validation.Get("age"))
	assert.NotEmpty(t, validation.Get("email"))
	assert.Equal(t, "The email field must be a valid email address", validation.Get("email"))
	assert.NotEmpty(t, validation.Get("password"))
	assert.Equal(t, "The password field must be minimum 6 char", validation.Get("password"))
	assert.NotEmpty(t, validation.Get("repeat_password"))
	assert.Equal(t, "The repeat_password field must be minimum 6 char", validation.Get("repeat_password"))
}

func TestUserUpdateRequestValidationOk(t *testing.T) {
	request := UserUpdateRequest{
		Name:    "Henrique",
		Age:     33,
		Address: "Rua Bla, 1234",
	}

	validation := request.ValidateUserUpdateRequest()

	assert.Empty(t, validation.Get("name"))
	assert.Empty(t, validation.Get("age"))
	assert.Empty(t, validation.Get("address"))
}

func TestUserUpdateRequestValidationWithEmptyValues(t *testing.T) {
	request := UserUpdateRequest{
		Name:    "",
		Address: "",
	}

	validation := request.ValidateUserUpdateRequest()

	assert.NotEmpty(t, validation.Get("name"))
	assert.Equal(t, "The name field is required", validation.Get("name"))
	assert.NotEmpty(t, validation.Get("age"))
	assert.Equal(t, "The age field is required", validation.Get("age"))
}

func TestUserUpdateRequestValidationWithInvalidValues(t *testing.T) {
	request := UserUpdateRequest{
		Name:    "He",
		Address: "",
	}

	validation := request.ValidateUserUpdateRequest()

	assert.NotEmpty(t, validation.Get("name"))
	assert.Equal(t, "The name field must be minimum 3 char", validation.Get("name"))
}
