# Data Models

## TODO

Fields:
- id: integer, primary key
- title: text, required
- notes: text, optional
- completed: boolean, default false
- created_at: timestamp
- updated_at: timestamp

## Storage Notes
- SQLite file stored locally (default `data/todos.db`).
- Keep SQL portable for PostgreSQL migration.
