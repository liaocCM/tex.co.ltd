package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Notes     string `json:"notes,omitempty"`
	Completed bool   `json:"completed"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateTodoRequest struct {
	Title string `json:"title"`
	Notes string `json:"notes"`
}

type UpdateTodoRequest struct {
	Title     string `json:"title"`
	Notes     string `json:"notes"`
	Completed bool   `json:"completed"`
}

const uidHeader = "X-User-Uid"

func registerRoutes(r *gin.Engine, db *sql.DB) {
	api := r.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api.GET("/todos", func(c *gin.Context) {
		uid, ok := requireUID(c)
		if !ok {
			return
		}
		status := strings.ToLower(c.DefaultQuery("status", "all"))
		todos, err := listTodos(db, uid, status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"todos": todos})
	})

	api.POST("/todos", func(c *gin.Context) {
		uid, ok := requireUID(c)
		if !ok {
			return
		}
		var req CreateTodoRequest
		if err := c.ShouldBindJSON(&req); err != nil || strings.TrimSpace(req.Title) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		item, err := createTodo(db, uid, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		c.JSON(http.StatusCreated, item)
	})

	api.GET("/todos/:id", func(c *gin.Context) {
		uid, ok := requireUID(c)
		if !ok {
			return
		}
		item, err := getTodo(db, uid, c.Param("id"))
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		c.JSON(http.StatusOK, item)
	})

	api.PUT("/todos/:id", func(c *gin.Context) {
		uid, ok := requireUID(c)
		if !ok {
			return
		}
		var req UpdateTodoRequest
		if err := c.ShouldBindJSON(&req); err != nil || strings.TrimSpace(req.Title) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		item, err := updateTodo(db, uid, c.Param("id"), req)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		c.JSON(http.StatusOK, item)
	})

	api.DELETE("/todos/:id", func(c *gin.Context) {
		uid, ok := requireUID(c)
		if !ok {
			return
		}
		err := deleteTodo(db, uid, c.Param("id"))
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		c.Status(http.StatusNoContent)
	})
}

func requireUID(c *gin.Context) (string, bool) {
	uid := strings.TrimSpace(c.GetHeader(uidHeader))
	if uid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return "", false
	}
	return uid, true
}
