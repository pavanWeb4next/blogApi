package handlers

import (
	"blog-api/database"
	"blog-api/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var blogID uint

func setupTestApp() *fiber.App {
	database.Connect()
	app := fiber.New()

	app.Post("/api/blog-post", CreateBlog)
	app.Get("/api/blog-post", GetBlogs)
	app.Get("/api/blog-post/:id", GetBlog)
	app.Patch("/api/blog-post/:id", UpdateBlog)
	app.Delete("/api/blog-post/:id", DeleteBlog)

	return app
}

func TestCreateBlog(t *testing.T) {
	app := setupTestApp()

	body := models.BlogPost{
		Title:       "Unit Test Blog",
		Description: "Testing Create",
		Body:        "Test body content",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/api/blog-post", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	respBody, _ := io.ReadAll(resp.Body)
	var created models.BlogPost
	json.Unmarshal(respBody, &created)
	assert.Equal(t, "Unit Test Blog", created.Title)

	blogID = created.ID 
}

func TestGetBlogs(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest("GET", "/api/blog-post", nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetBlog(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest("GET", "/api/blog-post/"+strconv.Itoa(int(blogID)), nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	respBody, _ := io.ReadAll(resp.Body)
	var blog models.BlogPost
	json.Unmarshal(respBody, &blog)
	assert.Equal(t, blogID, blog.ID)
}

func TestUpdateBlog(t *testing.T) {
	app := setupTestApp()

	updateData := models.BlogPost{
		Title:       "Updated Title",
		Description: "Updated Desc",
		Body:        "Updated Body",
	}
	jsonBody, _ := json.Marshal(updateData)

	req := httptest.NewRequest("PATCH", "/api/blog-post/"+strconv.Itoa(int(blogID)), bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	respBody, _ := io.ReadAll(resp.Body)
	var updated models.BlogPost
	json.Unmarshal(respBody, &updated)
	assert.Equal(t, "Updated Title", updated.Title)
}

func TestDeleteBlog(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest("DELETE", "/api/blog-post/"+strconv.Itoa(int(blogID)), nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
