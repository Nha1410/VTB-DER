# CLAUDE.md - VTB-DER

Gateway của TVBT gửi/nhận điện nghiệp vụ với **VSD** cho bù trừ, thanh toán chứng khoán phái sinh, theo
**Phụ lục 06 (QĐ 115/QĐ-VSD)**. Vai trò của service: nhận dữ liệu nghiệp vụ, **mã hoá thành điện FIN**
(ký tự ISO 15022, chuyển đổi tiếng Việt, đóng khung block 1/2/4/5) để gửi đi, và giải mã chiều ngược lại.
Một Go module (`github.com/Nha1410/VTB-DER`, Go 1.25), viết theo **clean architecture**.

Đây là guidance luôn được nạp cho Claude Code. Dùng làm **default, không phải checklist** bắt buộc áp vào
mọi thay đổi. Slash commands ở `.claude/commands/`, skills ở `.claude/skills/`, tài liệu spec đã chuyển
thể ở `docs/` (đọc `docs/README.md` trước, **không đọc lại file Word**), plan ở `plan/`.

> "Mã hoá" ở đây là **encode theo chuẩn điện của VSD** (khung FIN + bộ ký tự + chuyển đổi tiếng Việt).
> Phụ lục 06 không định nghĩa mã hoá mật mã cho nội dung điện; phiên làm việc và session key do GateWay
> Client của VSD lo. Đừng thêm lớp mã hoá nào khác khi chưa được chủ hệ thống yêu cầu.

## Layout & dependency rule

Dependency chỉ hướng vào trong: `model` → `service` → `handler` / `repository` + `client` → `common`.

| Package | Được import | Cấm import |
|---|---|---|
| `internal/model` | stdlib, `constant` | mọi package khác của repo |
| `internal/constant` | stdlib | mọi package khác của repo |
| `internal/common/*` | stdlib + thư viện ngoài | `model`, `service`, `handler` |
| `internal/service` | `model`, `constant`, `common/*` | `handler`, `repository`, `client`, `config` |
| `internal/repository`, `internal/client/*` | `model`, `constant`, `service`, `common/*` | `handler` |
| `internal/handler/*`, `internal/middleware` | `service`, `model`, `constant`, `common/*`, `config` | `repository`, `client` |
| entry point (`cmd/*`, `internal/server`) | tất cả - nơi **duy nhất** được wiring | - |

Hiện có:

| Package | Nội dung |
|---|---|
| `internal/constant/fin.go` | Hằng số của khung FIN (BIC VSD, hậu tố địa chỉ, giới hạn session/ISN, trailer) |
| `internal/model` | `FINMessage`, `CustomError` + mã `ErrBadRequest`, `ErrInternal` |
| `internal/common/iso15022` | Bộ ký tự hợp lệ + `Encode`/`Decode` tiếng Việt theo Bảng 2 (test đối chiếu từng dòng với `docs/05-bang-chuyen-doi-ky-tu.md`) |
| `internal/service/fin_encoder.go` | `EncodeFIN`: dựng block 1/2/4/5 và kiểm tra hợp lệ |

- `service` không biết transport (HTTP, file, TCP); lỗi ra khỏi service là `*model.CustomError`.
- `common/iso15022` là kỹ thuật thuần, **không** biết MT nào; chuyển đổi chỉ áp cho **giá trị field**, không áp
  cho cấu trúc dòng (`:`, `//`).
- Interface (port) khai báo ở phía **dùng nó**, mỗi flow một file trong `service/`.
- Import prefix: `github.com/Nha1410/VTB-DER/...`.

## Go backend defaults (always)

- Bám theo package layout, naming, imports, và dependency style đang có trước khi tạo convention mới.
- Giữ thay đổi nhỏ, gọn trong phạm vi yêu cầu; tránh refactor lan man hay viết lại kiến trúc.
- Giữ ranh giới tầng, nhưng **đừng thêm tầng/abstraction mới** chỉ vì "cho đúng sách". Một interface chỉ sinh ra khi có người thứ hai cần nó hoặc để test.
- Constructor-based dependency injection cho dependency thật; tránh global cho state/clients/config. Wiring chỉ ở entry point.
- **Ngoại lệ duy nhất là logger**: dùng thẳng `log/slog`, log ra stdout. Field có thể là dữ liệu nhà đầu tư
  (tên, số giấy tờ, địa chỉ, SĐT, email, số tài khoản) log ở **Debug**; Info/Error chỉ chứa method, loại điện,
  số hiệu tham chiếu, kết quả, lỗi hệ thống không mang payload. *(Mặc định thận trọng, chờ chủ hệ thống chốt.)*
- **Mọi hàm/method (và method trong interface) đặt tên rõ ràng cho tham số VÀ giá trị trả về** (named parameters + named returns). Ví dụ: `func EncodeFIN(msg model.FINMessage) (fin []byte, err error)`. Tên phải nói lên ý nghĩa, không dùng `r`/`s`/`x` mơ hồ.
- Code và GoDoc bằng **English**; tiếng Việt chỉ ở `docs/`, `plan/`, README.
- Propagate `context.Context` qua I/O, external call, và request-scoped work.
- Trả error tường minh và wrap bằng `fmt.Errorf("context: %w", err)` khi thêm ngữ cảnh hữu ích.
- Đóng resource bằng `defer`; hủy goroutine tường minh khi goroutine có thể sống lâu hơn caller.
- Validate/sanitize dữ liệu tại trust boundary: HTTP input, file trong thư mục `receive` của GateWay Client, env/config.
- Thêm timeout cho external call. Chỉ thêm retry / circuit breaker / auth / rate limit khi failure mode hoặc ranh giới service yêu cầu, hoặc khi user yêu cầu rõ.
- Ưu tiên stdlib. Chỉ thêm dependency mới khi nó loại bỏ complexity đáng kể (hiện **không có** dependency ngoài).
- Giữ comment hữu ích: giải thích quyết định không hiển nhiên, không mô tả lại code.

## VSD FIN specifics (bắt buộc khi dựng/đọc điện)

Chi tiết và nguồn ở `.cursor/rules/vsd-fin-integration.mdc` và `docs/`. Tóm tắt:

- Khung điện: `{1:F01<BIC8>AXXX<session4><ISN6>}{2:I<MT3><BIC8>XXXX<ưu tiên>[<giám sát>][<020>]}{4:…-}{5:{MAC:00000000}{CHK:F1DBCA886BBF}{TNG:}}`.
  **ISN không được trùng trong cùng một phiên.**
- BICCODE TVBT: `VSD` + 3 ký tự tên viết tắt + `XX`; VSD/CCP: `VSDSVN06`.
- Ưu tiên `U` ↔ giám sát `1`/`3`; `N` ↔ `2` hoặc bỏ trống.
- Ký tự hợp lệ: `a-z A-Z 0-9 / - ? : ( ) . , ' +`, space, CR/LF. Còn lại (tiếng Việt, `& # % \`, `/`) bên gửi đổi
  thành `?mã telex?`, bên nhận đổi ngược. **Độ dài field tính sau khi chuyển đổi.** `@` viết là `(at)`; `?` thô
  không gửi được.
- Số hiệu tham chiếu của yêu cầu (tag `20` / `20C::SEME`) phải duy nhất; phản hồi liên kết qua `20C::RELA` / tag `21`.
- NAK: tag `451`=1, mã lỗi ở tag `405` (`docs/04-bang-ma-loi.md`).
- Spec có chỗ mơ hồ/mâu thuẫn: đọc `docs/07-luu-y-va-diem-chua-nhat-quan.md` trước khi cài đặt một điện, và
  **hỏi** thay vì đoán khi điểm đó đang được liệt kê là chưa rõ.

## Testing (áp dụng cho file `*_test.go`)

- Ưu tiên table-driven test; `t.Run` với tên case mô tả; tên test `Test{Symbol}_{Scenario}`.
- Chỉ dùng `t.Parallel()` khi test không chia sẻ mutable state, env var, port, hay global.
- Không gọi mạng/thư mục thật trong unit test (`t.TempDir()`, `httptest.Server`, fake qua interface nhỏ).
- Cái gì suy ra từ bảng trong spec thì phải có test đọc bảng đã chuyển thể trong `docs/` và so từng dòng.
- Không commit BICCODE/tài khoản/dữ liệu nhà đầu tư/chứng thư/token thật; fixture là dữ liệu bịa.
- Assert hành vi và contract, không phải chi tiết cài đặt. Với bug fix, thêm test sẽ fail trước khi fix khi khả thi.

## Verification (trước khi tuyên bố "done")

Chạy thật trong session, đọc exit code và output - không đoán:

```bash
gofmt -l .
go build ./...
go test ./...
go vet ./...
```

Không nói "tests pass", "fixed", "done" mà chưa có output thực tế.
