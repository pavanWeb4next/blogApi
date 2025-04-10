package main

import (
	"blog-api/database"
	_ "blog-api/docs" // for Swagger
	"blog-api/routes"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Blog API
// @version 1.0
// @description CRUD Blog API with Go-Fiber
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	database.Connect()
	routes.Setup(app)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Listen(":3000")
}
