# Backend (Go + Gin)

## Setup

- Requires Go 1.21+
- SQLite file is stored at `data/todos.db` by default.

## Run

```bash
go run .
```

Environment variables:
- `TODO_DB_PATH` to override SQLite path.

## Notes
- Schema is created automatically on startup.
- Backend creates the SQLite directory if it does not exist.
- Keep SQL portable for PostgreSQL migration.
