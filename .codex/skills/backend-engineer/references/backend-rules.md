# Backend Rules

## Service structure

- Handlers for transport, services for business logic
- Repository or storage layer for persistence
- Keep request/response models separate from DB models

## Validation and errors

- Validate inputs at handler boundary
- Use consistent error codes and messages
- Map errors to HTTP status per contract

## Logging

- Include request ID and user ID when available
- Log at info for success, warn for recoverable, error for failures

## Testing guidance

- Add unit tests for services when logic is non-trivial
- Add handler tests for critical endpoints

## README notes

- Include build, run, and migration steps if any
- Record required env vars
