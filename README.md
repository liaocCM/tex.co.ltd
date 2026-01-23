# tex.co.ltd
A role-based agent skills system that simulates a real software company.

The system uses a CTO-centric orchestration model where specialized skills collaborate through structured artifacts (specs, architecture, design, code, tests, deployment), stored in shared memory across the software development lifecycle.

Communication is intentionally constrained to reduce noise, prevent spec drift, and maximize execution quality.

## Roles, jobs, and writable files

| Role | Job | Files may modify |
| --- | --- | --- |
| CTO Orchestrator | Own SDLC planning, governance decisions, role boundaries, and cross-skill coordination. | `00-governance/`, `01-requirements/user-stories.yaml` |
| Product Manager | Define scope, assumptions, user stories, and NFRs in requirements artifacts. | `01-requirements/` |
| System Architect | Translate requirements into system design, API contracts, and data models. | `02-architecture/` |
| UI Designer | Create interaction flows and ASCII wireframes aligned to requirements and data contracts. | `03-design/` |
| Frontend Engineer | Implement UI in TypeScript + React from design and API contracts. | `04-implementation/frontend/` |
| Backend Engineer | Implement Go + Gin services based on architecture and data models. | `04-implementation/backend/` |
| QA Engineer | Plan, execute, and document tests mapped to requirements. | `05-testing/` |
| DevOps Engineer | Produce deployment plans, runbooks, and release notes from implementation and QA signals. | `06-deployment/` |
