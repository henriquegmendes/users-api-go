package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"henrique.mendes/users-api/controller"
	"henrique.mendes/users-api/middlewares"
	"henrique.mendes/users-api/repository"
	"henrique.mendes/users-api/service"
)

func SetupRoutes(app *fiber.App) {
	authMiddleware := middlewares.NewAuthMiddleware(
		repository.NewUsersRepository(),
	)
	authController := controller.NewAuthController(
		service.NewUsersService(
			repository.NewUsersRepository(),
		),
	)
	usersController := controller.NewUsersController(
		service.NewUsersService(
			repository.NewUsersRepository(),
		),
	)

	app.Post("/users", authController.CreateUser)
	app.Post("/users/auth", authController.AuthUser)

	app.Use(
		jwtware.New(jwtware.Config{
			SigningKey: []byte("secret"),
		}),
		authMiddleware.CheckAuthUserExists,
	)

	app.Get("/users", usersController.GetUsers)
	app.Get("/users/:id", usersController.GetUserById)
	app.Put("/users", usersController.UpdateUser)
	app.Delete("/users", usersController.DeleteUser)
}
