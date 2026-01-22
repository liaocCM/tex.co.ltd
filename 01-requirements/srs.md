# Software Requirements Specification (SRS)

## Scope
Deliver a TODO application MVP for a single user with local persistence. The MVP supports creating, viewing, updating, completing, deleting, and filtering TODO items.

## Assumptions
- Single-user, local environment.
- Local persistence via SQLite file; migration to PostgreSQL is planned later.
- No authentication or multi-user support in MVP.

## Functional Requirements

FR-1 Create TODOs
- Allow users to create a TODO with a title and optional notes.

FR-2 List TODOs
- Display a list of TODOs with completion state and last updated timestamp.

FR-3 Update TODOs
- Allow editing title and notes.

FR-4 Complete/Reopen TODOs
- Allow toggling completion state; completed items are visually distinct.

FR-5 Delete TODOs
- Allow deleting a TODO.

FR-6 Filter TODOs
- Allow filtering by all/active/completed.

## Constraints
- Storage must be local (SQLite) and isolated to ease PostgreSQL migration.
- Avoid SQLite-specific SQL features to keep portability.
