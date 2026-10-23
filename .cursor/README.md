# Cursor setup (Go)

Rules, commands và skill cho repo `VTB-DER` khi làm việc bằng Cursor. Bản song song cho Claude Code
nằm ở `.claude/` + `CLAUDE.md`; hai bộ cùng nội dung, không xung đột.

## Rules

| File | Phạm vi | Nội dung |
|---|---|---|
| `rules/go-backend.mdc` | `alwaysApply: true` | Default Go backend + layout repo này |
| `rules/vsd-fin-integration.mdc` | `alwaysApply: true` | Ràng buộc cứng từ Phụ lục 06 (khung FIN, BIC, ký tự, tiếng Việt, tham chiếu duy nhất) |
| `rules/go-api.mdc` | glob `**/*.go` | Hướng dẫn HTTP layer (chưa có) |
| `rules/go-testing.mdc` | glob `**/*_test.go` | Table-driven, so byte-chính-xác với ví dụ trong `docs/` |

## Commands

- `/brainstorm` - design trước khi code
- `/write-plan` - viết implementation plan vào `plan/`
- `/execute-plan` - chạy plan kèm verification
- `/verify` - bằng chứng trước khi nói "done"
- `/debug` - debug có hệ thống

## Skills

- `skills/karpathy-guidelines` - làm rõ yêu cầu, chốt scope và định nghĩa "done"
