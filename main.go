package main

import (
	"github.com/gofiber/fiber/v2"
	"henrique.mendes/users-api/database"
	"henrique.mendes/users-api/migrations"
	"henrique.mendes/users-api/routes"
)

// @title Go Users Api Swagger Doc
// @version 1.0
// @description Swagger Documentation for Go Test
// @termsOfService http://swagger.io/terms/
// @BasePath /
func main() {
	database.Connect()
	migrations.CreateTables(database.DB)

	app := fiber.New()
	routes.SetupRoutes(app)
	routes.SwaggerRoutes(app)

	app.Listen(":3000")
}
