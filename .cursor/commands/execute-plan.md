---
description: Execute an implementation plan step by step with verification
---

# Execute plan

## Start

1. Read the plan the user provides (default: newest file in `plan/`).
2. Review critically - raise blockers before coding.
3. Create a todo list from plan tasks (use the TodoWrite tool).

## Per task

1. Mark task in progress.
2. Follow steps exactly; minimal scope.
3. Run verification commands from the plan (or `go build ./... && go test ./... && go vet ./...`).
4. Mark task complete only after verification succeeds.

## Stop and ask when

- Instructions are unclear or contradictory.
- Tests fail repeatedly after a focused fix attempt.
- Plan requires breaking changes the user did not approve.
- The plan conflicts with a spec constraint in `.cursor/rules/vsd-fin-integration.mdc`.

## Finish

- Summarize what changed and evidence (test output, build exit code).
- Do not claim success without fresh command output in this session.
