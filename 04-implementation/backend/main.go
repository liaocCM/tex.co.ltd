package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

type Todo struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Notes     string    `json:"notes,omitempty"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func main() {
	dbPath := os.Getenv("TODO_DB_PATH")
	if dbPath == "" {
		dbPath = "data/todos.db"
	}

	if err := ensureDBDir(dbPath); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := initSchema(db); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	api := r.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api.GET("/todos", func(c *gin.Context) {
		status := strings.ToLower(c.DefaultQuery("status", "all"))
		todos, err := listTodos(db, status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"todos": todos})
	})

	api.POST("/todos", func(c *gin.Context) {
		var req CreateTodoRequest
		if err := c.ShouldBindJSON(&req); err != nil || strings.TrimSpace(req.Title) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		item, err := createTodo(db, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		c.JSON(http.StatusCreated, item)
	})

	api.GET("/todos/:id", func(c *gin.Context) {
		item, err := getTodo(db, c.Param("id"))
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
		var req UpdateTodoRequest
		if err := c.ShouldBindJSON(&req); err != nil || strings.TrimSpace(req.Title) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		item, err := updateTodo(db, c.Param("id"), req)
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
		err := deleteTodo(db, c.Param("id"))
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

	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}

func initSchema(db *sql.DB) error {
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS todos (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  notes TEXT,
  completed INTEGER NOT NULL DEFAULT 0,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);
`)
	return err
}

func listTodos(db *sql.DB, status string) ([]Todo, error) {
	query := "SELECT id, title, notes, completed, created_at, updated_at FROM todos"
	args := []any{}
	if status == "active" {
		query += " WHERE completed = 0"
	} else if status == "completed" {
		query += " WHERE completed = 1"
	}
	query += " ORDER BY updated_at DESC"

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		var completed int
		var created string
		var updated string
		if err := rows.Scan(&t.ID, &t.Title, &t.Notes, &completed, &created, &updated); err != nil {
			return nil, err
		}
		t.Completed = completed == 1
		t.CreatedAt, _ = time.Parse(time.RFC3339, created)
		t.UpdatedAt, _ = time.Parse(time.RFC3339, updated)
		todos = append(todos, t)
	}
	return todos, rows.Err()
}

func createTodo(db *sql.DB, req CreateTodoRequest) (Todo, error) {
	now := time.Now().UTC()
	res, err := db.Exec(
		"INSERT INTO todos (title, notes, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		req.Title, req.Notes, 0, now.Format(time.RFC3339), now.Format(time.RFC3339),
	)
	if err != nil {
		return Todo{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return Todo{}, err
	}
	return getTodoByID(db, id)
}

func getTodo(db *sql.DB, id string) (Todo, error) {
	return getTodoByQuery(db, "SELECT id, title, notes, completed, created_at, updated_at FROM todos WHERE id = ?", id)
}

func getTodoByID(db *sql.DB, id int64) (Todo, error) {
	return getTodoByQuery(db, "SELECT id, title, notes, completed, created_at, updated_at FROM todos WHERE id = ?", id)
}

func getTodoByQuery(db *sql.DB, query string, arg any) (Todo, error) {
	var t Todo
	var completed int
	var created string
	var updated string
	row := db.QueryRow(query, arg)
	if err := row.Scan(&t.ID, &t.Title, &t.Notes, &completed, &created, &updated); err != nil {
		return Todo{}, err
	}
	t.Completed = completed == 1
	t.CreatedAt, _ = time.Parse(time.RFC3339, created)
	t.UpdatedAt, _ = time.Parse(time.RFC3339, updated)
	return t, nil
}

func updateTodo(db *sql.DB, id string, req UpdateTodoRequest) (Todo, error) {
	now := time.Now().UTC()
	res, err := db.Exec(
		"UPDATE todos SET title = ?, notes = ?, completed = ?, updated_at = ? WHERE id = ?",
		req.Title, req.Notes, boolToInt(req.Completed), now.Format(time.RFC3339), id,
	)
	if err != nil {
		return Todo{}, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return Todo{}, err
	}
	if affected == 0 {
		return Todo{}, sql.ErrNoRows
	}
	return getTodo(db, id)
}

func deleteTodo(db *sql.DB, id string) error {
	res, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func boolToInt(val bool) int {
	if val {
		return 1
	}
	return 0
}

func ensureDBDir(dbPath string) error {
	dir := filepath.Dir(dbPath)
	if dir == "." || dir == "" {
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}
