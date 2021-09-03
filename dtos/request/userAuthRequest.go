package request

import (
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

type UserAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserAuthRequest) ValidateUserAuthRequest() url.Values {
	rules := govalidator.MapData{
		"email":    []string{"required", "min:4", "max:100", "email"},
		"password": []string{"required", "min:6", "max:100"},
	}

	opts := govalidator.Options{
		Data:  &u,
		Rules: rules,
	}
	v := govalidator.New(opts)

	return v.ValidateStruct()
}
