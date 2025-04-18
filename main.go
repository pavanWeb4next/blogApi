package main

import (
	_ "blog-api/docs" // for Swagger
	"blog-api/pkg/config"
	"blog-api/pkg/database"
	"blog-api/pkg/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Blog API
// @version 1.0
// @description CRUD Blog API with Go-Fiber
// @host
// @BasePath /
// @schemes https
func main() {
	config.LoadConfig()
	fmt.Println("Using DB:", config.AppConfig.DBType)
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	database.Connect()
	routes.Setup(app)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Listen(":8000")
}
