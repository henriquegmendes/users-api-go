package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"henrique.mendes/users-api/dtos/request"
	"henrique.mendes/users-api/service"
)

func NewAuthController(service *service.UsersService) *UsersController {
	return &UsersController{
		service: service,
	}
}

// Register User.
// @Description Register a new User
// @Summary Register a new User
// @Tags Public Routes
// @Produce json
// @Param data body request.UserCreateRequest true "User Data"
// @Success 200 {object} response.UserResponse
// @Router /users/register [post]
func (contr UsersController) CreateUser(c *fiber.Ctx) error {
	data := new(request.UserCreateRequest)
	if err := c.BodyParser(data); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}

	validation := data.ValidateUserCreateRequest()
	if len(validation) != 0 {
		return c.Status(400).JSON(validation)
	}

	userResponse, err := contr.service.Create(data)
	if err != nil {
		return c.Status(400).Send([]byte(err.Error()))
	}

	return c.Status(201).JSON(userResponse)
}

// Authenticate User
// @Description Authenticate User Based on email/password Credentials
// @Summary Authenticate User Based on email/password Credentials
// @Tags Public Routes
// @Produce json
// @Param data body request.userAuthRequest true "User Data"
// @Success 200 {object} response.userAuthResponse
// @Router /users/auth [post]
func (contr UsersController) AuthUser(c *fiber.Ctx) error {
	data := new(request.UserAuthRequest)
	if err := c.BodyParser(data); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}

	validation := data.ValidateUserAuthRequest()
	if len(validation) != 0 {
		return c.Status(400).JSON(validation)
	}

	userResponse, err := contr.service.FindByEmail(*data)
	if err != nil {
		if strings.Contains(err.Error(), "credentials") {
			return c.Status(400).JSON(fiber.Map{
				"message": err.Error(),
			})
		} else {
			return c.Status(503).Send([]byte(err.Error()))
		}
	}

	return c.Status(200).JSON(userResponse)
}
