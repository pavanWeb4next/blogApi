package models

import "time"

type BlogPost struct {
	ID          uint      `json:"id" example:"1"`
	Title       string    `json:"title" example:"new blog"`
	Description string    `json:"description" example:"my personal blog"`
	Body        string    `json:"body" example:"personal details"`
	CreatedAt   time.Time `json:"created_at" example:"2025-04-11T01:15:30.123Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2025-04-11T01:20:45.456Z"`
}

type BlogPostInput struct {
	Title       string `json:"title" example:"Hello Fiber"`
	Description string `json:"description" example:"Getting started with Go-Fiber"`
	Body        string `json:"body" example:"This is a blog body"`
}
