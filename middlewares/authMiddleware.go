package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"henrique.mendes/users-api/repository"
	"henrique.mendes/users-api/utils"
)

func CheckAuthUserExists(c *fiber.Ctx) error {
	userId, error := utils.GetTokenInfo(c)
	if error != nil {
		c.Status(503).JSON(error.Error())
		return nil
	}

	user := repository.FindById(uint(userId))
	if user.Id == 0 {
		c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized user",
		})
		return nil
	}

	c.Next()
	return nil
}
