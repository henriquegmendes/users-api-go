package request

import (
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

type UserCreateRequest struct {
	Name           string `json:"name"`
	Age            int    `json:"age"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
	Address        string `json:"address"`
}

func (u UserCreateRequest) ValidateUserCreateRequest() url.Values {
	rules := govalidator.MapData{
		"name":            []string{"required", "min:3", "max:100"},
		"age":             []string{"required"},
		"email":           []string{"required", "min:4", "max:100", "email"},
		"password":        []string{"required", "min:6", "max:100"},
		"repeat_password": []string{"required", "min:6", "max:100"},
		"address":         []string{"min:3", "max:200"},
	}

	opts := govalidator.Options{
		Data:  &u,
		Rules: rules,
	}
	v := govalidator.New(opts)

	return v.ValidateStruct()
}
