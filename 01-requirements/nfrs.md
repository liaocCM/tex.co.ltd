# Non-Functional Requirements (NFRs)

## Performance
- List retrieval should complete within 500ms for up to 1,000 TODOs on a local machine.
- Create/update/delete should complete within 300ms locally.

## Reliability
- Data must persist across restarts.
- On startup, the system should recover with no manual steps.

## Security
- No authentication in MVP; limit exposure to localhost by default.

## Maintainability
- Separate data access from handlers to ease storage migration.
- Keep API contracts stable and versioned if changed.
