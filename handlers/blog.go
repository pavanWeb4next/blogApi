package handlers

import (
	"blog-api/database"
	"blog-api/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary Create a blog post
// @Tags Blog
// @Accept json
// @Produce json
// @Param blog body models.BlogPost true "Blog Post"
// @Success 201 {object} models.BlogPost
// @Failure 400 {object} map[string]interface{}
// @Router /api/blog-post [post]
func CreateBlog(c *fiber.Ctx) error {
	blog := new(models.BlogPost)
	if err := c.BodyParser(blog); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	database.DB.Create(&blog)
	return c.Status(201).JSON(blog)
}

// @Summary Get all blog posts
// @Tags Blog
// @Produce json
// @Success 200 {array} models.BlogPost
// @Failure 404 {object} map[string]interface{}
// @Router /api/blog-post [get]
func GetBlogs(c *fiber.Ctx) error {
	var blogs []models.BlogPost
	database.DB.Find(&blogs)
	return c.JSON(blogs)
}

// @Summary Get a single blog post by ID
// @Tags Blog
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200 {object} models.BlogPost
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/blog-post/{id} [get]
func GetBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	var blog models.BlogPost
	result := database.DB.First(&blog, id)
	if result.Error == gorm.ErrRecordNotFound {
		return c.Status(404).JSON(fiber.Map{"error": "Blog not found"})
	}
	return c.JSON(blog)
}

// @Summary Update a blog post
// @Tags Blog
// @Accept json
// @Produce json
// @Param id path int true "Blog ID"
// @Param blog body models.BlogPost true "Blog Post"
// @Success 200 {object} models.BlogPost
// @Failure 404 {object} map[string]interface{}
// @Router /api/blog-post/{id} [patch]
func UpdateBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	var blog models.BlogPost
	result := database.DB.First(&blog, id)
	if result.Error == gorm.ErrRecordNotFound {
		return c.Status(404).JSON(fiber.Map{"error": "Blog not found"})
	}

	updateData := new(models.BlogPost)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	database.DB.Model(&blog).Updates(updateData)
	return c.JSON(blog)
}

// @Summary Delete a blog post
// @Tags Blog
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/blog-post/{id} [delete]
func DeleteBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	var blog models.BlogPost
	result := database.DB.First(&blog, id)
	if result.Error == gorm.ErrRecordNotFound {
		return c.Status(404).JSON(fiber.Map{"error": "Blog not found"})
	}

	database.DB.Delete(&blog)
	return c.JSON(fiber.Map{"message": "Deleted"})
}
