---
description: Explore intent and design before any implementation (Superpowers-style brainstorm)
---

# Brainstorm

Do not write or change production code in this session.

## Steps

1. Explore the current project (layout, `go.mod`, `internal/`, existing patterns).
2. Read the relevant part of `docs/` (start at `docs/README.md`) before proposing anything that touches the FIN format or a business flow.
3. Ask clarifying questions **one at a time** - purpose, constraints, success criteria.
4. If the idea spans several independent services (account registration vs margin vs settlement vs reports), propose splitting into smaller deliverables first.
5. Propose **2-3 approaches** with trade-offs and a clear recommendation.
6. Present a short design (architecture, main components, data flow, errors, testing). Ask for approval after each major section.
7. When design is approved, summarize in a short spec saved under `plan/`.

## Hard gate

No implementation, scaffolding, or dependency changes until the user approves the design.

## After approval

Offer to produce an implementation plan (use the **write-plan** command) or stop if the user only wanted design.
