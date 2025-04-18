package routes

import (
	"blog-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	blog := api.Group("/blog-post")
	blog.Post("/", handlers.CreateBlog)
	blog.Get("/", handlers.GetBlogs)
	blog.Get("/:id", handlers.GetBlog)
	blog.Patch("/:id", handlers.UpdateBlog)
	blog.Delete("/:id", handlers.DeleteBlog)
}
