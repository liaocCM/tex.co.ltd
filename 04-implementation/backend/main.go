package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

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
	registerRoutes(r, db)

	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}

func ensureDBDir(dbPath string) error {
	dir := filepath.Dir(dbPath)
	if dir == "." || dir == "" {
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}
