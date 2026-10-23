# Backend Standards & Best Practices

Chuẩn kỹ thuật cho toàn bộ code Go trong repo `VTB-DER`. Port từ repo `VNEID` (cùng chủ, cùng phong cách),
bỏ phần đặc thù VNeID/Gin/Oracle. Kết hợp với nguyên tắc Karpathy ở skill `.claude/skills/karpathy-guidelines/`.
`CLAUDE.md` là bản rút gọn luôn được nạp; file này là bản đầy đủ.

## Go General Guidelines

### Basic Principles

- **Dùng English** cho code và GoDoc comment (tiếng Việt chỉ dùng trong `docs/`, `plan/`, `README`).
- **Khai báo type tường minh** cho tham số và giá trị trả về.
- **Named parameters + named returns** cho mọi hàm/method, kể cả method trong interface.
- **Tránh `interface{}` / `any`** trừ khi thật cần (ngoại lệ chấp nhận: `args ...any` của hàm kiểu `fmt.Sprintf`).
- **GoDoc cho mọi symbol exported.**
- **`gofmt -w .`** trước khi commit.

### Nomenclature

- **PascalCase** cho exported: `FINMessage`, `EncodeFIN`.
- **camelCase** cho biến, hàm nội bộ: `receiverBIC`, `validateFrame`.
- **snake_case** cho tên file: `fin_encoder.go`.
- **UPPERCASE** cho env var.
- Không magic number/string: `constant.MaxSessionNumber`, `constant.VSDBIC`.
- Boolean prefix `is`/`has`/`can`/`should`: `isValid`, `insideMarker`.
- Viết đủ chữ, trừ viết tắt chuẩn: `API`, `URL`, `ID`, `BIC`, `ISN`, `MT`, `ctx`, `cfg`, `msg`, `err`.

### Functions

- Ngắn: < 20 dòng (lý tưởng < 10). Một mức abstraction cho mỗi hàm.
- Verb + object: `EncodeFIN`, `ValidateCharset`, `validatePriority`.
- Giảm lồng nhau bằng early return.
- Nhiều input → gom vào struct (`model.FINMessage`), đừng truyền 6 tham số rời.

### Error Handling

- Idiom Go: `if err != nil`. Không `panic()` trong service code.
- Wrap có ngữ cảnh: `fmt.Errorf("decode text: %w", err)`.
- **Không bao giờ bỏ qua error** bằng `_` trần (trừ `defer x.Close()` theo idiom).
- Lỗi ra khỏi `service` là `*model.CustomError` với `Code` từ `model.Err*`; chi tiết nội bộ chỉ vào log.

---

## Kiến trúc - clean architecture

Service là cổng hai chiều giữa nghiệp vụ của TVBT và VSD: chiều đi **dựng điện FIN**, chiều về **giải điện**.

| Vòng | Package | Trách nhiệm |
|---|---|---|
| Entity | `internal/model` | Kiểu nghiệp vụ thuần (`FINMessage`), `CustomError`. **Zero dependency** ngoài `constant`. |
| Constant | `internal/constant` | Giá trị cố định của chuẩn điện (BIC VSD, hậu tố, giới hạn, trailer). |
| Use case | `internal/service` | Dựng/kiểm tra điện (`EncodeFIN`); sau này: dựng từng loại MT, giải điện, ACK/NAK. |
| Interface adapter | `internal/repository`, `internal/client/*`, `internal/handler/*` | (chưa có) ghi/đọc thư mục GateWay Client, HTTP... |
| Infrastructure | `internal/common/*` | Kỹ thuật thuần: `iso15022` (bộ ký tự, chuyển đổi tiếng Việt). |
| Wiring | entry point | Nơi **duy nhất** khởi tạo và nối dependency. |

**Chiều phụ thuộc chỉ hướng vào trong.** Bảng cấm/cho phép import: xem `CLAUDE.md`.

### Nguyên tắc

- **Handler mỏng:** bind → gọi service qua interface → map response.
- **Interface khai báo ở phía dùng nó**, không ở phía implement.
- **Usecase độc lập framework.**
- **Lỗi qua biên là `*model.CustomError`.**
- **Thời gian tiêm vào** qua `Clock` interface khi cần (chưa cần ở v0).
- **Constructor DI:** mọi dependency truyền vào `New*`; không global cho state. Ngoại lệ: logger (`log/slog`).

### Anti-patterns

- ❌ Nghiệp vụ trong handler
- ❌ `model`/`service` import framework HTTP
- ❌ `common` import ngược lên `model`/`service`
- ❌ Wiring rải rác ngoài entry point
- ❌ Interface một-implement chỉ để "cho đúng clean architecture"
- ❌ Chuyển đổi ký tự lên **cấu trúc** dòng điện thay vì chỉ giá trị field
- ❌ Tính độ dài field **trước** khi chuyển đổi tiếng Việt
- ❌ Gửi điện khi chưa kiểm bộ ký tự (để VSD NAK thay vì bắt lỗi sớm)
- ❌ Dùng lại ISN trong cùng một phiên

---

## Testing Standards

- Unit test cho `model`, `service`, `common`, mục tiêu > 80%.
- Table-driven + `t.Run` với tên case mô tả; Arrange-Act-Assert.
- So sánh **byte-chính-xác** với ví dụ trong `docs/` khi có.
- Test suy từ bảng spec: đọc bảng trong `docs/` và so từng dòng (xem `TestEncode_MatchesSpecTable`).
- Tên test: `Test{Symbol}_{Scenario}`; biến: `inputMsg`, `wantFIN`, `gotFIN`.

```go
func TestValidateCharset(t *testing.T) {
    tests := []struct {
        name      string
        inputText string
        wantIndex int
    }{
        {name: "brace would break the block structure", inputText: "A}", wantIndex: 1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) { /* ... */ })
    }
}
```

---

## Configuration & Environment

- 12-factor: mọi cấu hình qua biến môi trường. **Không commit secret** (chứng thư, token BKAV, mật khẩu GateWay Client).
- Validate lúc startup; cấu hình sai thì **fail startup**, không chạy tiếp.

---

## Code Quality

```bash
gofmt -w .
go build ./...
go test ./...
go vet ./...
golangci-lint run ./...   # nếu đã cài; cấu hình ở .golangci.yml
```

Linter bật: `errcheck`, `govet`, `staticcheck`, `unused`, `ineffassign`, `gosimple`, `misspell`.

---

## Summary

1. **Code quality:** gofmt, lint, type tường minh, named returns
2. **Clean code:** hàm nhỏ, tên nói lên ý nghĩa, một trách nhiệm
3. **Kiến trúc:** dependency hướng vào trong, wiring chỉ ở entry point
4. **Testing:** table-driven, byte-chính-xác với ví dụ spec, đối chiếu bảng spec từng dòng
5. **Đúng chuẩn VSD:** khung FIN, bộ ký tự, chuyển đổi trước khi đếm độ dài, ISN duy nhất
6. **Config/secret:** env, validate lúc startup, không secret trong code
