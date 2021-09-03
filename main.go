package main

import (
	"github.com/gofiber/fiber/v2"
	"henrique.mendes/users-api/database"
	"henrique.mendes/users-api/migrations"
	"henrique.mendes/users-api/routes"
)

func main() {
	database.Connect()
	migrations.CreateTables(database.DB)

	app := fiber.New()
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
