# 001 - Khởi tạo dự án và mã hoá điện FIN

**Goal:** module Go 1.25 có thể dựng điện FIN gửi lên VSD đúng Phụ lục 06, gồm chuyển đổi tiếng Việt.

**Architecture:** clean architecture 4 vòng như repo `VNEID`. Kỹ thuật thuần (bộ ký tự, telex) ở
`common/iso15022`; khung điện ở `service`; kiểu dữ liệu ở `model`/`constant`. Không có dependency ngoài.

## Context

Vai trò của dự án là gateway "mã hoá" thông tin gửi đi. Phụ lục 06 không có mã hoá mật mã cho nội dung
điện: phần bảo mật (đăng nhập, session key, chữ ký số BKAV) do GateWay Client của VSD xử lý. Phần thuộc
gateway là **encode theo chuẩn điện**: khung FIN block 1/2/4/5, bộ ký tự hợp lệ và chuyển đổi tiếng Việt
sang ISO 15022. Đó là phạm vi plan này.

## Đã làm

- [x] `go mod init github.com/Nha1410/VTB-DER` (Go 1.25)
- [x] `internal/constant/fin.go` - hằng số khung điện
- [x] `internal/model` - `FINMessage`, `CustomError`
- [x] `internal/common/iso15022` - `IsValidChar`, `ValidateCharset`, `Encode`, `Decode`; test đối chiếu
  139 dòng Bảng 2 với `docs/05-bang-chuyen-doi-ky-tu.md` và hai ví dụ VD1/VD2 của spec
- [x] `internal/service/fin_encoder.go` - `EncodeFIN` + kiểm tra block 1/2 và bộ ký tự block 4
- [x] Port rule/skill/command từ `VNEID` (`.claude/`, `.cursor/`, `CLAUDE.md`, `backend-standards.md`)

## Quyết định đã chốt (có thể đổi)

1. **Ngắt dòng block 4 là CRLF** - spec không nêu; theo chuẩn SWIFT FIN.
2. **`?` thô trong text bị từ chối** (không gửi được) vì bên nhận sẽ đọc thành dấu mở marker; spec liệt `?` là
   ký tự hợp lệ nhưng không nói cách thoát.
3. **`_` chỉ hợp lệ bên trong cặp `?…?`** - Bảng 2 dùng `?_?`, `?_38?`… nhưng bộ ký tự mục 2.1 không có `_`.
4. **`@` bị từ chối khi `Encode`** - spec bảo viết `(at)` cho email; việc đó thuộc tầng dựng từng field, chưa làm.
5. **Chuyển đổi chỉ áp lên giá trị field**, không áp lên cấu trúc dòng; caller gọi `Encode` trên giá trị trước khi ghép dòng.
6. **Chuỗi vào phải ở dạng NFC** (chữ có dấu dựng sẵn).
7. **ISN do caller giữ** - chưa có sequencer.

## Việc tiếp theo (đề xuất, chưa duyệt)

- [ ] Sequencer ISN theo phiên (gateway phải đảm bảo không trùng) - cần chốt nơi lưu trạng thái
- [ ] Builder từng loại điện, bắt đầu nhóm 1 (MT598 600/602) rồi các nhóm khác - `docs/nghiep-vu/`
- [ ] `FINDecoder`: tách block 1/2/4, nhận diện ACK/NAK (`451`/`405`), `Decode` tiếng Việt chiều về
- [ ] Lớp I/O: ghi `.fin` vào `send`, quét `receive`/`error` của GateWay Client
- [ ] Xác nhận với VSD các điểm trong `docs/07-luu-y-va-diem-chua-nhat-quan.md` (CRLF, `?`, `_`, định dạng tag 177)

## Verification

```bash
gofmt -l .
go build ./...
go test ./...
go vet ./...
```
