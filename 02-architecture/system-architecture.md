# System Architecture

## Overview
The MVP is a simple client-server application.

- Frontend: React + TypeScript SPA
- Backend: Go + Gin REST API
- Storage: SQLite file (local), data access isolated for PostgreSQL migration

## Components

### Frontend
- Renders TODO list and controls.
- Calls REST API for CRUD and filtering.

### Backend API
- Provides REST endpoints for TODO operations.
- Validates input and returns JSON responses.

### Data Access
- SQLite-backed repository layer.
- Initializes schema on startup.

## Deployment Shape
- Local development: frontend and backend run separately.
- Backend binds to localhost by default; frontend uses proxy or API base URL.
