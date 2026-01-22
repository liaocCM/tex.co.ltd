import React, { useEffect, useState } from "react";
import { createTodo, deleteTodo, listTodos, updateTodo } from "./api";

export type Todo = {
  id: number;
  title: string;
  notes?: string;
  completed: boolean;
  created_at: string;
  updated_at: string;
};

type Filter = "all" | "active" | "completed";

const emptyForm = { title: "", notes: "" };

export default function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [filter, setFilter] = useState<Filter>("all");
  const [form, setForm] = useState(emptyForm);
  const [editing, setEditing] = useState<Todo | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    void refresh();
  }, [filter]);

  async function refresh() {
    setLoading(true);
    setError(null);
    try {
      const data = await listTodos(filter);
      setTodos(data);
    } catch (err) {
      setError("Failed to load todos");
    } finally {
      setLoading(false);
    }
  }

  async function handleCreate(event: React.FormEvent) {
    event.preventDefault();
    if (!form.title.trim()) {
      setError("Title is required");
      return;
    }
    setError(null);
    try {
      await createTodo(form.title.trim(), form.notes.trim());
      setForm(emptyForm);
      await refresh();
    } catch (err) {
      setError("Failed to create todo");
    }
  }

  async function handleUpdate(event: React.FormEvent) {
    event.preventDefault();
    if (!editing || !form.title.trim()) {
      setError("Title is required");
      return;
    }
    setError(null);
    try {
      await updateTodo(editing.id, {
        title: form.title.trim(),
        notes: form.notes.trim(),
        completed: editing.completed
      });
      setEditing(null);
      setForm(emptyForm);
      await refresh();
    } catch (err) {
      setError("Failed to update todo");
    }
  }

  async function handleToggle(todo: Todo) {
    setError(null);
    try {
      await updateTodo(todo.id, {
        title: todo.title,
        notes: todo.notes ?? "",
        completed: !todo.completed
      });
      await refresh();
    } catch (err) {
      setError("Failed to update todo");
    }
  }

  async function handleDelete(id: number) {
    setError(null);
    try {
      await deleteTodo(id);
      await refresh();
    } catch (err) {
      setError("Failed to delete todo");
    }
  }

  function startEdit(todo: Todo) {
    setEditing(todo);
    setForm({ title: todo.title, notes: todo.notes ?? "" });
  }

  function cancelEdit() {
    setEditing(null);
    setForm(emptyForm);
  }

  return (
    <div className="app">
      <header className="app-header">
        <div>
          <p className="eyebrow">Anthropic TODO</p>
          <h1>TODO MVP</h1>
          <p className="subtle">A focused list with a calm, editorial tone.</p>
        </div>
        <div className="filter-buttons">
          <button
            className={`filter-button ${filter === "all" ? "is-active" : ""}`}
            onClick={() => setFilter("all")}
          >
            All
          </button>
          <button
            className={`filter-button ${filter === "active" ? "is-active" : ""}`}
            onClick={() => setFilter("active")}
          >
            Active
          </button>
          <button
            className={`filter-button ${filter === "completed" ? "is-active" : ""}`}
            onClick={() => setFilter("completed")}
          >
            Completed
          </button>
        </div>
      </header>

      {error && (
        <p style={{ color: "crimson" }}>{error}</p>
      )}

      <section className="panel">
        <div className="panel-header">
          <h2>{editing ? "Edit TODO" : "Add TODO"}</h2>
          <span className="tag">Notes optional</span>
        </div>
        <form onSubmit={editing ? handleUpdate : handleCreate} className="form">
          <div className="form-field">
            <label htmlFor="todo-title">Title</label>
            <input
              id="todo-title"
              type="text"
              placeholder="Give this TODO a clear title"
              value={form.title}
              onChange={(event) => setForm({ ...form, title: event.target.value })}
            />
          </div>
          <div className="form-field">
            <label htmlFor="todo-notes">Notes</label>
            <textarea
              id="todo-notes"
              placeholder="Add optional detail"
              value={form.notes}
              onChange={(event) => setForm({ ...form, notes: event.target.value })}
            />
          </div>
          <div className="form-actions">
            <button className="primary" type="submit">
              {editing ? "Save" : "Add"}
            </button>
            {editing && (
              <button className="ghost" type="button" onClick={cancelEdit}>
                Cancel
              </button>
            )}
          </div>
        </form>
      </section>

      <section className="panel">
        <div className="panel-header">
          <h2>List</h2>
          <span className="tag">{todos.length} items</span>
        </div>
        {loading && <p>Loading...</p>}
        {!loading && todos.length === 0 && <p>No TODOs yet.</p>}
        <ul className="todo-list">
          {todos.map((todo) => (
            <li key={todo.id} className="todo-item">
              <div className="todo-row">
                <label className="todo-title">
                  <input
                    type="checkbox"
                    checked={todo.completed}
                    onChange={() => handleToggle(todo)}
                  />
                  <span className={todo.completed ? "is-complete" : ""}>{todo.title}</span>
                </label>
                <div className="todo-actions">
                  <button className="ghost" onClick={() => startEdit(todo)}>Edit</button>
                  <button className="danger" onClick={() => handleDelete(todo.id)}>Delete</button>
                </div>
              </div>
              {todo.notes && <p className="todo-notes">{todo.notes}</p>}
              <small className="todo-meta">Updated: {new Date(todo.updated_at).toLocaleString()}</small>
            </li>
          ))}
        </ul>
      </section>
    </div>
  );
}
