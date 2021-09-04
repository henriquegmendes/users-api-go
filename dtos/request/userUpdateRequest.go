package request

import (
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

type UserUpdateRequest struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func (u UserUpdateRequest) ValidateUserUpdateRequest() url.Values {
	rules := govalidator.MapData{
		"name": []string{"required", "min:3", "max:100"},
		"age":  []string{"required"},
	}

	opts := govalidator.Options{
		Data:  &u,
		Rules: rules,
	}
	v := govalidator.New(opts)

	return v.ValidateStruct()
}
