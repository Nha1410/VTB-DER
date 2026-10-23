---
description: Create a bite-sized implementation plan from an approved design or spec
---

# Write plan

Create an implementation plan. Assume the reader has little context about this repo.

## Before tasks

- Map files to create or modify and each file's responsibility.
- Follow existing project structure and conventions (`internal/{model,constant,service,common}`; wiring only at the entry point once one exists).
- Quote the spec constraints the plan must respect (block layout, BICCODE rules, character set and conversion, unique ISN per session) and cite the `docs/` file they come from.
- Each task should be independently verifiable where possible.

## Plan format

```markdown
# [Feature] - implementation plan

**Goal:** one sentence

**Architecture:** 2-3 sentences

**Tech stack:** from go.mod / existing code

---

### Task 1: [name]

**Files:** list paths

- [ ] Step 1: …
- [ ] Step 2: …
- [ ] Verify: command or check
```

## Task size

Each step should be one small action (roughly 2-5 minutes): failing test → implement → run tests → commit (if user asked for commits).

## Rules

- DRY, YAGNI; TDD when the user wants tests first.
- No placeholders or TODOs in planned code snippets.
- Include exact verification commands (`go test ./...`, `go vet ./...`, curl examples).

Save the plan as `plan/NNN-<slug>.md` (next free number), or output it in chat if the user prefers.
