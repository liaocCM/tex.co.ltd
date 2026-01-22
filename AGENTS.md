# Agent Organization Specification (AGENT.md)

## Overview

This repository defines a **CTO-orchestrated, role-based agent organization** that simulates a real software company.

Agents (skills) collaborate through **explicit artifacts**—requirements, architecture, design, code, tests, and deployment documents—stored in a **layered shared memory** aligned with the Software Development Life Cycle (SDLC).

Communication is intentionally constrained to reduce noise, prevent specification drift, and maximize execution quality. The CTO agent acts as the single authority for orchestration, approvals, and conflict resolution.

---

## Core Principles

1. **Artifacts over conversation**  
   All meaningful work is written to files. Chat is never treated as a source of truth.

2. **Layered shared memory**  
   Each SDLC stage has its own directory. Agents may read upstream layers but may only write to their own layer.

3. **CTO-centric orchestration**  
   The CTO agent plans work, approves changes, and enforces collaboration boundaries.

4. **Minimal but intentional communication**  
   Communication occurs only via structured governance artifacts (RFC / CR).

---

## Agent Roles

| Agent | Responsibility |
|-----|---------------|
| CTO | Orchestration, planning, final decision authority |
| PM | Software requirements and user stories |
| Architect | System architecture, APIs, data models |
| UI Designer | ASCII-based UI and interaction design |
| Frontend Engineer | TypeScript + React implementation |
| Backend Engineer | Golang + Gin implementation |
| QA Engineer | Test planning, execution, and reporting |
| DevOps Engineer | Deployment and operational readiness |

---

## Shared Memory Structure

All agents share the `memory/` directory:

```
memory/
├── 00-governance/        # Rules, decisions, approvals
├── 01-requirements/     # SRS, user stories, NFRs
├── 02-architecture/     # System design, APIs, data models
├── 03-design/           # UI / UX (ASCII)
├── 04-implementation/   # Frontend & backend code
├── 05-testing/          # Test plans, cases, results
└── 06-deployment/       # Deployment plans & runbooks
```

### Global Rules

- Agents **read upstream layers only**
- Agents **write only to their assigned layer**
- Upstream changes require a **Change Request (CR)** and CTO approval
- All major decisions must be recorded

---

## Governance Artifacts

### Decision Log

`memory/00-governance/decision-log.md`

Records:
- Architectural choices
- Scope decisions
- Conflict resolutions

### Change Requests (CR)

`memory/00-governance/change-requests.md`

Used to propose:
- Requirement changes
- Architecture changes
- Cross-layer updates

### Collaboration Rules

`memory/00-governance/collaboration-rules.md`

Defines:
- Who may communicate
- About which topics
- Under what conditions

---

## Communication Protocol

### Request for Clarification (RFC)

Used when an agent needs clarification without proposing a change.

```
## RFC: <title>
From: <Agent>
To: <Agent>
Reason:
Questions:
Impact if unanswered:
CTO Approval Needed: Yes / No
```

### Change Request (CR)

Used when proposing a modification to an existing artifact.

```
## CR: <title>
Requested by:
Affected file:
Proposed change:
Reason:
Risk:
Status: Proposed | Approved | Rejected
CTO Notes:
```

---

## Authority Model

### CTO Authority

The CTO agent:
- Breaks down user requests into an execution plan
- Assigns work to agents
- Approves or rejects Change Requests
- Controls collaboration boundaries
- Ensures consistency across SDLC layers

**CTO decisions are final.**

### Allowed Direct Interaction (Controlled)

- **PM ↔ Architect**  
  Allowed only via RFC documents and limited to requirement feasibility and constraints.

- **Frontend ↔ Backend**  
  Allowed only via API contract Change Requests.

All other communication requires explicit CTO approval.

---

## Execution Flow (Typical)

1. User submits a request
2. CTO creates an SDLC execution plan
3. PM produces requirements
4. Architect designs the system
5. UI Designer creates design documents
6. Engineers implement code
7. QA validates behavior
8. DevOps prepares and executes deployment
9. CTO reviews and approves release

---

## Success Criteria

The system is successful when:
- Artifacts remain consistent across layers
- Late-stage rework is minimized
- Decisions are traceable
- Execution converges reliably

---

## Non-Goals

This system does **not** aim to:
- Simulate human conversation
- Encourage free-form brainstorming
- Optimize creativity at the expense of correctness

The system is designed to optimize for **clarity, control, and delivery**.

---

## Final Note

This agent organization is intentionally opinionated.

If agents disagree:
- They write it down
- They escalate via governance artifacts
- The CTO decides

This constraint is a feature, not a limitation.