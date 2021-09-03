package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"henrique.mendes/users-api/controller"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/users", controller.Create)
	app.Post("/auth", controller.Auth)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	app.Get("/teste", controller.Test)
}
