package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"henrique.mendes/users-api/dtos/request"
	"henrique.mendes/users-api/service"
	"henrique.mendes/users-api/utils"
)

func GetUsers(c *fiber.Ctx) error {
	name := c.Query("name", "")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	response := service.FindByNamePaginated(name, page, limit)

	return c.JSON(response)
}

func GetUserById(c *fiber.Ctx) error {
	userId, error := strconv.ParseUint(c.Params("id"), 10, 64)
	if error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Id param must be a number",
		})
	}

	response := service.FindById(uint(userId))
	if response.Id == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(response)
}

func UpdateUser(c *fiber.Ctx) error {
	userId, error := utils.GetTokenInfo(c)
	if error != nil {
		return c.Status(503).JSON(error.Error())
	}

	data := new(request.UserUpdateRequest)
	if err := c.BodyParser(data); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}

	response := service.UpdateUserById(userId, *data)
	if response.Id == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(response)
}

func DeleteUser(c *fiber.Ctx) error {
	userId, error := utils.GetTokenInfo(c)
	if error != nil {
		return c.Status(503).JSON(error.Error())
	}

	if error := service.DeleteUserById(userId); error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": error.Error(),
		})
	}

	return c.SendStatus(204)
}
