---
description: Verify claims with evidence before saying work is complete
---

# Verify

Before stating that work is done, fixed, or passing:

1. **Identify** the command that proves the claim.
2. **Run** the full command in this session (not a previous run).
3. **Read** exit code and full output.
4. **Report** actual status with evidence.

## Go defaults (module root is the repo root - `go.mod` next to `README.md`)

```bash
go build ./...
go test ./...
go vet ./...
```

Add `golangci-lint run ./...` or project-specific commands if the user or plan requires them.
The module has no runnable binary yet, so the unit tests are the runtime check. Inside the Mirae network
`go` commands may need `GOPROXY=direct`.

## Forbidden without evidence

- "Tests pass", "build succeeded", "fixed", "done", "looks good"
- Words like should, probably, seems

If verification fails, state what failed and what was attempted - do not guess the fix.
