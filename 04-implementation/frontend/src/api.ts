import type { Todo } from "./App";

type TodoResponse = {
  todos: Todo[];
};

type CreateTodo = {
  title: string;
  notes?: string;
};

type UpdateTodo = {
  title: string;
  notes?: string;
  completed: boolean;
};

const baseUrl = "/api";
const uidStorageKey = "todo_uid";

function getUid(): string {
  if (typeof window === "undefined") {
    return "server";
  }
  const existing = window.localStorage.getItem(uidStorageKey);
  if (existing && existing.trim()) {
    return existing;
  }
  const newUid =
    typeof window.crypto?.randomUUID === "function"
      ? window.crypto.randomUUID()
      : `uid-${Date.now()}-${Math.random().toString(16).slice(2)}`;
  window.localStorage.setItem(uidStorageKey, newUid);
  return newUid;
}

async function request<T>(path: string, options?: RequestInit): Promise<T> {
  const response = await fetch(`${baseUrl}${path}`, {
    headers: {
      "Content-Type": "application/json",
      "X-User-Uid": getUid()
    },
    ...options
  });
  if (!response.ok) {
    throw new Error(`Request failed: ${response.status}`);
  }
  if (response.status === 204) {
    return undefined as T;
  }
  return response.json() as Promise<T>;
}

export async function listTodos(status: string): Promise<Todo[]> {
  const data = await request<TodoResponse | null>(`/todos?status=${status}`);
  return data?.todos ?? [];
}

export async function createTodo(title: string, notes: string): Promise<Todo> {
  const payload: CreateTodo = { title, notes: notes || undefined };
  return request<Todo>("/todos", {
    method: "POST",
    body: JSON.stringify(payload)
  });
}

export async function updateTodo(id: number, payload: UpdateTodo): Promise<Todo> {
  return request<Todo>(`/todos/${id}`, {
    method: "PUT",
    body: JSON.stringify(payload)
  });
}

export async function deleteTodo(id: number): Promise<void> {
  await request<void>(`/todos/${id}`, { method: "DELETE" });
}
