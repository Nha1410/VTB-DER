---
name: karpathy-guidelines
description: Use when the user's request is vague, too broad, or before starting implementation - sharpen the question, clarify scope, and define done criteria (Karpathy principles).
---

# Karpathy guidelines - sharpen the request

Use this skill to help the **user** ask better questions and to align on scope **before** writing code.

## When to use

- Request is ambiguous ("làm nốt phần mã hoá đi", "fix cái ký tự tiếng Việt", "thêm điện MT598")
- Multiple valid interpretations exist
- No clear definition of done
- User wants a plan or review without a large diff yet

## 1. Think before coding

- State what is unclear; do not guess silently.
- Offer 2-3 interpretations: "Do you mean A, B, or C?"
- Ask one focused question at a time when possible.
- List assumptions explicitly and ask the user to confirm or reject.
- For FIN work, name which MT message and which spec section (`docs/`) the request maps to.

## 2. Simplicity first

- Restate the **minimum** scope that satisfies the stated goal.
- Call out extras the user did not ask for (new abstractions, extra endpoints, speculative config,
  a DB layer before it is needed).
- Prefer validating at boundaries (HTTP input, VSD responses, files in `receive`) over internal defensive layers that cannot fail.

## 3. Surgical changes (for implementation requests)

- Confirm: "Only touch X; leave Y unchanged?"
- Warn if the request implies unrelated refactors or style sweeps.
- Prefer matching existing project patterns over introducing new ones.

## 4. Goal-driven execution

Turn the request into verifiable criteria, for example:

- "Done when `go test ./...` passes with no failures."
- "Done when `EncodeFIN` returns `ErrBadRequest` for a message that still holds unconverted Vietnamese, and the framed bytes match the spec example."
- "Done when the Vietnamese encoder matches all 139 rows of Bảng 2 (test reads `docs/05-bang-chuyen-doi-ky-tu.md`)."

Ask: **How will we know this is finished?**

## Prompt coaching (help the user write better asks)

| Element | Example |
|--------|---------|
| Goal | "Thêm dựng điện MT598 mở tài khoản" |
| Constraints | "Không thêm dependency; giữ contract API nội bộ" |
| Location | "Chỉ trong `internal/service`" |
| Done | "`go test ./...` xanh; điện dựng ra khớp ví dụ trong `docs/02-cau-truc-dien-fin.md`" |
| Non-goals | "Chưa làm HTTP, chưa làm sequencer ISN" |

## Output format

1. **Restated goal** (one sentence)
2. **Open questions** (numbered)
3. **Proposed scope** (in / out)
4. **Definition of done** (checklist)
5. **Recommended next step** (e.g. small spike, test first, or implement)

Do not start implementation until the user confirms scope or explicitly says to proceed with stated assumptions.
