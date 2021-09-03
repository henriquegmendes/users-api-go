package routes

import (
	"github.com/gofiber/fiber/v2"
	"henrique.mendes/users-api/controller"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/users", controller.Create)
	app.Post("/auth", controller.Auth)
}
