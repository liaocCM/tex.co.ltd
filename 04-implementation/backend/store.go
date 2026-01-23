package main

import (
	"database/sql"
	"fmt"
	"time"
)

func initSchema(db *sql.DB) error {
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS todos (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uid TEXT NOT NULL DEFAULT 'legacy',
  title TEXT NOT NULL,
  notes TEXT,
  completed INTEGER NOT NULL DEFAULT 0,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);
`)
	if err != nil {
		return err
	}
	if err := ensureColumn(db, "todos", "uid", "TEXT NOT NULL DEFAULT 'legacy'"); err != nil {
		return err
	}
	if _, err := db.Exec("CREATE INDEX IF NOT EXISTS todos_uid_idx ON todos(uid)"); err != nil {
		return err
	}
	if _, err := db.Exec("CREATE INDEX IF NOT EXISTS todos_uid_updated_idx ON todos(uid, updated_at)"); err != nil {
		return err
	}
	return nil
}

func listTodos(db *sql.DB, uid, status string) ([]Todo, error) {
	query := "SELECT id, title, notes, completed, created_at, updated_at FROM todos WHERE uid = ?"
	args := []any{uid}
	if status == "active" {
		query += " AND completed = 0"
	} else if status == "completed" {
		query += " AND completed = 1"
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
		if err := rows.Scan(&t.ID, &t.Title, &t.Notes, &completed, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		t.Completed = completed == 1
		todos = append(todos, t)
	}
	return todos, rows.Err()
}

func createTodo(db *sql.DB, uid string, req CreateTodoRequest) (Todo, error) {
	now := time.Now().UTC().Format(time.RFC3339)
	res, err := db.Exec(
		"INSERT INTO todos (uid, title, notes, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		uid, req.Title, req.Notes, 0, now, now,
	)
	if err != nil {
		return Todo{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return Todo{}, err
	}
	return getTodoByID(db, uid, id)
}

func getTodo(db *sql.DB, uid, id string) (Todo, error) {
	return getTodoByQuery(db, "SELECT id, title, notes, completed, created_at, updated_at FROM todos WHERE uid = ? AND id = ?", uid, id)
}

func getTodoByID(db *sql.DB, uid string, id int64) (Todo, error) {
	return getTodoByQuery(db, "SELECT id, title, notes, completed, created_at, updated_at FROM todos WHERE uid = ? AND id = ?", uid, id)
}

func getTodoByQuery(db *sql.DB, query string, args ...any) (Todo, error) {
	var t Todo
	var completed int
	row := db.QueryRow(query, args...)
	if err := row.Scan(&t.ID, &t.Title, &t.Notes, &completed, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return Todo{}, err
	}
	t.Completed = completed == 1
	return t, nil
}

func updateTodo(db *sql.DB, uid, id string, req UpdateTodoRequest) (Todo, error) {
	now := time.Now().UTC().Format(time.RFC3339)
	res, err := db.Exec(
		"UPDATE todos SET title = ?, notes = ?, completed = ?, updated_at = ? WHERE uid = ? AND id = ?",
		req.Title, req.Notes, boolToInt(req.Completed), now, uid, id,
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
	return getTodo(db, uid, id)
}

func deleteTodo(db *sql.DB, uid, id string) error {
	res, err := db.Exec("DELETE FROM todos WHERE uid = ? AND id = ?", uid, id)
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

func ensureColumn(db *sql.DB, table, column, ddl string) error {
	rows, err := db.Query(fmt.Sprintf("PRAGMA table_info(%s)", table))
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cid int
		var name string
		var ctype string
		var notnull int
		var dflt sql.NullString
		var pk int
		if err := rows.Scan(&cid, &name, &ctype, &notnull, &dflt, &pk); err != nil {
			return err
		}
		if name == column {
			return nil
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	_, err = db.Exec(fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", table, column, ddl))
	return err
}
