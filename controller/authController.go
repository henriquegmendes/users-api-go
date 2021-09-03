package controller

import (
	"github.com/gofiber/fiber/v2"
	"henrique.mendes/users-api/dtos/request"
	"henrique.mendes/users-api/mappers"
	"henrique.mendes/users-api/service"
)

func CreateUser(c *fiber.Ctx) error {
	data := new(request.UserCreateRequest)
	if err := c.BodyParser(data); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}

	validation := data.ValidateUserCreateRequest()
	if len(validation) != 0 {
		return c.Status(400).JSON(validation)
	}

	user, err := service.Create(data)
	if err != nil {
		return c.Status(400).Send([]byte(err.Error()))
	}

	return c.Status(201).JSON(mappers.ToUserResponse(user))
}

func AuthUser(c *fiber.Ctx) error {
	data := new(request.UserAuthRequest)
	if err := c.BodyParser(data); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}

	validation := data.ValidateUserAuthRequest()
	if len(validation) != 0 {
		return c.Status(400).JSON(validation)
	}

	user := service.FindByEmail(data.Email)

	if user.Id == 0 || user.HasInvalidPassword(data.Password) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Wrong credentials",
		})
	}

	response, err := mappers.ToUserAuthResponse(user)
	if err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}

	return c.Status(200).JSON(response)
}
