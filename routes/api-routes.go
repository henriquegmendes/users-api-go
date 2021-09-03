package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"henrique.mendes/users-api/controller"
	"henrique.mendes/users-api/middlewares"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/users", controller.CreateUser)
	app.Post("/users/auth", controller.AuthUser)

	app.Use(
		jwtware.New(jwtware.Config{
			SigningKey: []byte("secret"),
		}),
		middlewares.CheckAuthUserExists,
	)

	app.Get("/users", controller.GetUsers)
	app.Get("/users/:id", controller.GetUserById)
	app.Put("/users", controller.UpdateUser)
	app.Delete("/users", controller.DeleteUser)
}
