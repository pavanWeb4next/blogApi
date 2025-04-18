package handlers

import (
	"blog-api/models"
	"blog-api/pkg/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary Create a blog post
// @Tags Blog
// @Accept json
// @Produce json
// @Param blog body models.BlogPostInput true "Blog Post"
// @Success 201 {object} models.BlogPost
// @Failure 400 {object} map[string]interface{}
// @Router /api/blog-post [post]
func CreateBlog(c *fiber.Ctx) error {
	var input models.BlogPostInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	blog := models.BlogPost{
		Title:       input.Title,
		Description: input.Description,
		Body:        input.Body,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
// @Param blog body models.BlogPostInput true "Blog Post"
// @Success 200 {object} models.BlogPost
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/blog-post/{id} [patch]
func UpdateBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	var blog models.BlogPost
	result := database.DB.First(&blog, id)
	if result.Error == gorm.ErrRecordNotFound {
		return c.Status(404).JSON(fiber.Map{"error": "Blog not found"})
	}

	var updateData map[string]interface{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Only update non-empty fields
	if title, ok := updateData["title"].(string); ok && title != "" {
		blog.Title = title
	}
	if desc, ok := updateData["description"].(string); ok && desc != "" {
		blog.Description = desc
	}
	if body, ok := updateData["body"].(string); ok && body != "" {
		blog.Body = body
	}

	blog.UpdatedAt = time.Now()
	database.DB.Save(&blog)
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
