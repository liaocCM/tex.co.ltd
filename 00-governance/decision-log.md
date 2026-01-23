# Decision Log

## Execution Plan (MVP: TODO App)

### Stage 1: Requirements
- Owner: Product Manager
- Deliverables: `01-requirements/srs.md`, `01-requirements/user-stories.yaml`, `01-requirements/nfrs.md`
- Notes: Define MVP scope and acceptance criteria for CRUD, list filtering, and basic UX.
- Status: completed

### Stage 2: Architecture
- Owner: System Architect
- Deliverables: `02-architecture/system-architecture.md`, `02-architecture/api-contracts.yaml`, `02-architecture/data-models.md`
- Notes: Service/API design, storage layer, migration considerations.
- Status: completed

### Stage 3: Design
- Owner: UI Designer
- Deliverables: `03-design/ui-ascii.md`, `03-design/user-flows.md`
- Notes: Low-fidelity wireframes for list/detail/edit flows.
- Status: completed

### Stage 4: Implementation
- Owners: Frontend Engineer, Backend Engineer
- Deliverables: `04-implementation/frontend/` and `04-implementation/backend/`
- Notes: React + TypeScript frontend, Go + Gin backend.
- Status: in-progress (code scaffolded)

### Stage 5: QA
- Owner: QA Engineer
- Deliverables: `05-testing/test-plan.md`, `05-testing/test-cases.md`, `05-testing/test-results.md`
- Notes: Map cases to user stories; record execution results.
- Status: in-progress (plan/cases drafted)

### Stage 6: Deployment
- Owner: DevOps Engineer
- Deliverables: `06-deployment/deployment-plan.md`, `06-deployment/runbook.md`, `06-deployment/release-notes.md`
- Notes: Local/dev deployment and ops guidance.
- Status: in-progress (docs drafted)

## Execution Plan (UID Data Isolation)

### Stage 1: Requirements Update
- Owner: Product Manager
- Deliverables: Update `01-requirements/srs.md` and `01-requirements/nfrs.md` to include `uid` scoping.
- Notes: Capture data isolation rules and localStorage `uid` dependency.
- Status: in-progress

### Stage 2: Architecture Update
- Owner: System Architect
- Deliverables: Update `02-architecture/api-contracts.yaml` and `02-architecture/data-models.md`.
- Notes: Add `uid` field and API requirements; define migration/compatibility.
- Status: planned

### Stage 3: Implementation Update
- Owners: Frontend Engineer, Backend Engineer
- Deliverables: Update `04-implementation/frontend/` and `04-implementation/backend/`.
- Notes: Frontend reads `uid` from localStorage and passes it; backend scopes all queries by `uid`.
- Status: planned

### Stage 4: QA Update
- Owner: QA Engineer
- Deliverables: Update `05-testing/test-cases.md` and `05-testing/test-results.md`.
- Notes: Add cross-uid visibility tests and localStorage switching coverage.
- Status: planned

## Incident: Backend Startup Failure

### Triage Summary
- Symptom: `go run .` fails with "unable to open database file".
- Likely cause: SQLite file path uses `data/todos.db` but `data/` is missing in backend working dir.

### Hotfix Plan
- Owner: Backend Engineer
- Deliverable: Ensure backend creates the DB directory or documents required setup.
- Status: in-progress

## Key Decisions

- Storage for MVP: Use SQLite with a local file (in-repo or under a `data/` directory). Keep data access isolated to ease migration to PostgreSQL.
- Migration intent: Plan for PostgreSQL by keeping SQL simple and avoiding SQLite-specific features.
- MVP focus: Single-user, local storage, core TODO CRUD and basic filtering.
- Branding: Apply Anthropic brand palette and typography for UI styling per `03-design/brand-guidelines.md`.
- Requirement update: Separate TODO records by client `uid` stored in localStorage, with all CRUD scoped to `uid`.

## Review Notes (UID Scope)

- Requirements review: `01-requirements/srs.md` and `01-requirements/nfrs.md` still reflect single-user scope; needs `uid` isolation updates.
- Architecture review: `02-architecture/api-contracts.yaml` and `02-architecture/data-models.md` do not include `uid` field or per-uid scoping.

## CTO Decision (Proceed Without Updated Artifacts)

- Proceed with implementation under provisional rules: `uid` is required on all TODO requests; server scopes all CRUD by `uid`; schema adds `uid` (text) with index; existing records default to a placeholder `uid` until migrated.

## Collaboration Boundaries

- PM and Architect may coordinate via RFC/CR artifacts only; no freeform chat.
- Engineers file CRs for requirements status updates and contract changes.

## Dispatches

- PM: Draft `01-requirements/srs.md` and `01-requirements/nfrs.md` from the MVP scope; refine `01-requirements/user-stories.yaml` if needed.
- Architect: Stand by for requirements; then produce `02-architecture/system-architecture.md`, `02-architecture/api-contracts.yaml`, `02-architecture/data-models.md`.
- UI Designer: Stand by for requirements and API contracts; then produce `03-design/ui-ascii.md` and `03-design/user-flows.md`.
- Frontend Engineer: Stand by for design + contracts; then implement `04-implementation/frontend/` and document `04-implementation/frontend/README.md`.
- Backend Engineer: Stand by for architecture + models; then implement `04-implementation/backend/` and document `04-implementation/backend/README.md`.
- QA Engineer: Stand by for implementation notes; then create `05-testing/test-plan.md`, `05-testing/test-cases.md`, `05-testing/test-results.md`.
- DevOps Engineer: Stand by for QA results; then create `06-deployment/deployment-plan.md`, `06-deployment/runbook.md`, `06-deployment/release-notes.md`.
- Architect: Update API contracts and data model to scope TODOs by `uid` from localStorage.
- Frontend Engineer: Read `uid` from localStorage and include it in TODO API requests.
- Backend Engineer: Scope queries/mutations by `uid` and enforce per-uid data isolation.
- PM: Update `01-requirements/srs.md` and `01-requirements/nfrs.md` to document `uid` scoping and data isolation.
- QA Engineer: Add `uid` isolation and localStorage switching cases once implementation updates land.
- Architect: Update `02-architecture/api-contracts.yaml` and `02-architecture/data-models.md` to include `uid` scoping and migration notes.
- Frontend Engineer: Implement `uid` scoping in API calls after updated API contracts land.
- Backend Engineer: Implement `uid` scoping in data access after updated data model/API contracts land.
- Frontend Engineer: Proceed now with provisional `uid` rules; add localStorage `uid` read + include `uid` in all TODO API requests.
- Backend Engineer: Proceed now with provisional `uid` rules; add `uid` field and enforce per-uid scoping in all queries/mutations.
