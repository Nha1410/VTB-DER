---
description: Systematic debugging - reproduce, isolate root cause, then fix
---

# Debug

Do not change code until the failure is understood.

## Process

1. **Reproduce** - exact steps, input, expected vs actual.
2. **Evidence** - logs, stack trace, failing test name, HTTP status/body.
3. **Hypothesis** - one likely root cause.
4. **Test hypothesis** - smallest check (unit test, log line, curl).
5. **Fix** - minimal change that addresses root cause only.
6. **Verify** - rerun failing test/command; check for regressions.

## Go habits

- Read the full error chain (`%w` wraps).
- Prefer a failing test that captures the bug before patching.
- For flaky tests: `go test -race ./...` when appropriate.

## VSD FIN-specific suspects

Check these before blaming our code (details in `docs/07-luu-y-va-diem-chua-nhat-quan.md`):

- A NAK from VSD: read tag `405` against `docs/04-bang-ma-loi.md` (`H..` header, `T..` text, `T98` duplicate).
- Duplicate ISN within one session, or a reused transaction reference (tag `20` / `20C::SEME`).
- Vietnamese, or `@ & # % \`, left unconverted; or a literal `?` in free text.
- Field length counted **before** conversion (the limit applies **after**, see `docs/03-ky-tu-va-chuyen-doi-tieng-viet.md`).
- Converting the structural characters of a line (`:`, `//`) instead of just the field value.
- Wrong BICCODE form: members are `VSD` + 3 letters + `XX`, VSD/CCP is `VSDSVN06`.
- Block 2: urgent priority without delivery monitoring 1/3, or normal with anything but 2/none.
- A file left in the GateWay Client `error` folder, or never picked up from `send`.

## Stop and ask

- Cannot reproduce.
- Fix would require large refactor or unclear requirements.

After fix, run `go test ./...` and `go vet ./...` unless the user specifies otherwise.
