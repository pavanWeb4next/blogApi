package handlers_test

import (
	"blog-api/handlers"
	"blog-api/models"
	"blog-api/pkg/config"
	"blog-api/pkg/database"
	"strconv"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupApp() *fiber.App {
	mockConfig()
	database.Connect()

	app := fiber.New()
	app.Post("/api/blog-post", handlers.CreateBlog)
	app.Get("/api/blog-post", handlers.GetBlogs)
	app.Get("/api/blog-post/:id", handlers.GetBlog)
	app.Patch("/api/blog-post/:id", handlers.UpdateBlog)
	app.Delete("/api/blog-post/:id", handlers.DeleteBlog)
	return app
}

func mockConfig() {
	config.AppConfig = config.Config{
		DBType:      "sqlite",
		PostgresDSN: "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable",
		SQLitePath:  "test_blog.db",
		MongoURI:    "mongodb://localhost:27017",
		MongoDBName: "blogdb",
		UseMemcache: false,
	}
}

func TestCreateBlog(t *testing.T) {
	app := setupApp()

	payload := `{
		"title": "Test Blog",
		"description": "Testing create blog",
		"body": "This is the body of the blog"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/blog-post", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, 201, resp.StatusCode)

	var blog models.BlogPost
	json.NewDecoder(resp.Body).Decode(&blog)
	assert.Equal(t, "Test Blog", blog.Title)
	assert.Equal(t, "Testing create blog", blog.Description)
	assert.Equal(t, "This is the body of the blog", blog.Body)
	assert.WithinDuration(t, time.Now(), blog.CreatedAt, time.Second*2)
}

func TestGetBlogs(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest(http.MethodGet, "/api/blog-post", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	var blogs []models.BlogPost
	err := json.NewDecoder(resp.Body).Decode(&blogs)
	assert.Nil(t, err)
}

func TestGetBlog_NotFound(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest(http.MethodGet, "/api/blog-post/999999", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 404, resp.StatusCode)
}

func TestUpdateBlog(t *testing.T) {
	app := setupApp()

	// Create a blog first
	blog := models.BlogPost{
		Title:       "Old Title",
		Description: "Old Desc",
		Body:        "Old Body",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	database.DB.Create(&blog)

	payload := `{
        "title": "Updated Title",
        "description": "Updated Desc",
        "body": "Updated Body"
    }`

	req := httptest.NewRequest(http.MethodPatch, "/api/blog-post/"+strconv.Itoa(int(blog.ID)), strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var updated models.BlogPost
	json.NewDecoder(resp.Body).Decode(&updated)
	assert.Equal(t, "Updated Title", updated.Title)
}

func TestDeleteBlog(t *testing.T) {
	app := setupApp()

	// Create a blog first
	blog := models.BlogPost{
		Title:       "Delete Me",
		Description: "To be deleted",
		Body:        "Bye!",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	database.DB.Create(&blog)

	req := httptest.NewRequest(http.MethodDelete, "/api/blog-post/"+strconv.Itoa(int(blog.ID)), nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var msg map[string]string
	json.NewDecoder(resp.Body).Decode(&msg)
	assert.Equal(t, "Deleted", msg["message"])
}
