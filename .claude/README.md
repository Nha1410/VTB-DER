# Claude Code setup (Go)

Bộ guidance cho repo `VTB-DER`, song song với `.cursor/`. Port từ repo `VNEID` (cùng chủ); phần đặc thù
VNeID AgentGW được thay bằng rule từ Phụ lục 06 của VSD.

## Ánh xạ từ `.cursor/`

| Cursor | Claude Code | Ghi chú |
|--------|-------------|---------|
| `rules/go-backend.mdc` (always-on) | `CLAUDE.md` (mục *Layout*, *Go backend defaults*) | Luôn nạp vào context |
| `rules/go-api.mdc` (glob `*.go`) | `backend-standards.md` | Chưa có HTTP layer; để dành |
| `rules/go-testing.mdc` (glob `*_test.go`) | `CLAUDE.md` (mục *Testing*) | Claude Code không auto-apply theo glob |
| `rules/vsd-fin-integration.mdc` (always-on) | `CLAUDE.md` (mục *VSD FIN specifics*) | Ràng buộc từ spec, không phải preference |
| `commands/*.md` | `.claude/commands/*.md` | Gọi bằng slash command, vd `/brainstorm` |
| `skills/karpathy-guidelines` | `.claude/skills/karpathy-guidelines/SKILL.md` | Format skill tương thích |

## Contents

**Commands** (slash commands, self-contained):

- `/brainstorm` - design before code
- `/write-plan` - implementation plan, lưu vào `plan/NNN-<slug>.md`
- `/execute-plan` - chạy plan kèm verification
- `/verify` - evidence trước khi nói "done"
- `/debug` - debug có hệ thống, có checklist nghi phạm riêng của điện FIN/VSD

**Skills**:

- `karpathy-guidelines` - làm rõ yêu cầu và định nghĩa "done" trước khi code

## Tài liệu liên quan

- `CLAUDE.md` - guidance rút gọn, luôn nạp mỗi session
- `backend-standards.md` - chuẩn kỹ thuật đầy đủ
- `docs/` - Phụ lục 06 đã chuyển thể sang markdown (đọc `docs/README.md`, không đọc lại file Word)
- `plan/` - implementation plan

## Notes

- Không mang `.claude/settings.json` / `settings.local.json` từ VNEID sang: đó là allowlist quyền riêng của
  máy/phiên làm việc (có đường dẫn cá nhân), không phải guidance. Tạo lại bằng `/fewer-permission-prompts`
  hoặc `/update-config` khi cần.
- Thư mục `.cursor/` giữ song song để ai dùng Cursor vẫn hoạt động; hai bộ không xung đột.
