# VTB-DER

Gateway của Thành viên bù trừ (TVBT) trao đổi điện nghiệp vụ với VSD cho bù trừ, thanh toán chứng khoán
phái sinh theo **Phụ lục 06 (QĐ 115/QĐ-VSD ngày 31/05/2017)**.

Go module `github.com/Nha1410/VTB-DER`, Go 1.25, clean architecture. Hiện chưa có dependency ngoài.

## Trạng thái

| Hạng mục | Trạng thái |
|---|---|
| Tài liệu spec chuyển thể markdown | Xong - [docs/README.md](docs/README.md) |
| Bộ ký tự + chuyển đổi tiếng Việt (`internal/common/iso15022`) | Xong, đối chiếu 139/139 dòng Bảng 2 |
| Dựng khung điện FIN block 1/2/4/5 (`service.EncodeFIN`) | Xong |
| Dựng từng loại MT (598, 103, 524, 542…) | Chưa làm |
| Giải điện chiều về + ACK/NAK | Chưa làm |
| Ghi/đọc thư mục GateWay Client, sequencer ISN, HTTP | Chưa làm |

Xem [plan/](plan/README.md) cho kế hoạch.

## Layout

```
internal/
  constant/             hằng số khung FIN
  model/                FINMessage, CustomError
  common/iso15022/      bộ ký tự hợp lệ, Encode/Decode tiếng Việt
  service/              EncodeFIN
docs/                   Phụ lục 06 đã chuyển thể + file Word gốc
plan/                   implementation plan
.claude/ .cursor/       rule, command, skill (CLAUDE.md, backend-standards.md ở gốc)
```

## Dùng thử

```go
name, _ := iso15022.Encode("NGUYỄN VĂN A")          // NGUY?EEX?N V?AW?N A
fin, err := service.EncodeFIN(model.FINMessage{
    SenderBIC: "VSDSSIXX", SessionNumber: 20, SequenceNumber: 21,
    MessageType: "598", Priority: "N",
    Text: []string{":20:REF0001", ":12:600", ":95Q::INVE//" + name},
})
```

## Kiểm tra

```bash
gofmt -l .
go build ./...
go test ./...
go vet ./...
```
