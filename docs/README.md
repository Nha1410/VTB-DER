# Tài liệu VTB-DER — Phụ lục 06 VSD (chuẩn điện bù trừ, thanh toán CKPS)

Bộ file markdown này là bản chuyển thể từ [Phu-luc-06.docx](Phu-luc-06.docx) để tra cứu mà không phải đọc lại file Word. Nguồn: **Phụ lục 06 ban hành kèm Quyết định 115/QĐ-VSD ngày 31/05/2017** của Tổng Giám đốc Trung tâm Lưu ký Chứng khoán Việt Nam — *Phương thức trao đổi và chuẩn điện nghiệp vụ qua Cổng giao tiếp trực tuyến của VSD cho hoạt động bù trừ, thanh toán giao dịch chứng khoán phái sinh*.

## Quy ước

- Các bảng đặc tả điện, bảng mã lỗi, bảng ký tự được chuyển **nguyên văn, đủ dòng, đủ cột** (cột `No.` giữ đúng số trong tài liệu gốc, kể cả chỗ bị lặp).
- Sơ đồ luồng trong file Word là ảnh nhúng; đã được dựng lại thành sơ đồ **Mermaid** từ chữ trong ảnh (kết hợp phần mô tả trình tự).
- Ký tự `<br>` trong bảng thay cho xuống dòng trong ô; `•` thay cho gạch đầu dòng trong ô.
- Phần **diễn giải/suy luận** (không phải nguyên văn) luôn được đánh dấu: [00](00-tong-quan-va-thuat-ngu.md) mục 3–5, [06](06-danh-muc-dien-mt.md), [08](08-checklist-trien-khai.md), phần "Quy luật rút ra" trong [03](03-ky-tu-va-chuyen-doi-tieng-viet.md).
- Lỗi/mâu thuẫn của văn bản gốc **không được sửa** trong các file nguyên văn; xem [07](07-luu-y-va-diem-chua-nhat-quan.md).

## Mục lục

| File | Nội dung |
|---|---|
| [00-tong-quan-va-thuat-ngu.md](00-tong-quan-va-thuat-ngu.md) | Phạm vi, các bên, từ viết tắt, quy ước tài khoản `BICCODE.C/P/R/S`, ký hiệu định dạng trường |
| [01-phuong-thuc-trao-doi.md](01-phuong-thuc-trao-doi.md) | GateWay Client, thư mục send/receive/archive/error, phiên làm việc, token BKAV |
| [02-cau-truc-dien-fin.md](02-cau-truc-dien-fin.md) | Block 1/2/4/5, BICCODE, ACK/NAK, ví dụ |
| [03-ky-tu-va-chuyen-doi-tieng-viet.md](03-ky-tu-va-chuyen-doi-tieng-viet.md) | Bộ ký tự hợp lệ, quy tắc `?..?`, cách tính độ dài, ví dụ |
| [04-bang-ma-loi.md](04-bang-ma-loi.md) | Bảng 1 — mã lỗi |
| [05-bang-chuyen-doi-ky-tu.md](05-bang-chuyen-doi-ky-tu.md) | Bảng 2 — chuyển đổi tiếng Việt/ký tự đặc biệt |
| [06-danh-muc-dien-mt.md](06-danh-muc-dien-mt.md) | Tra cứu nhanh: điện MT → mục → chiều → sub-type/`77E` |
| [07-luu-y-va-diem-chua-nhat-quan.md](07-luu-y-va-diem-chua-nhat-quan.md) | Lỗi/mâu thuẫn trong văn bản gốc, câu hỏi cần xác nhận với VSD |
| [08-checklist-trien-khai.md](08-checklist-trien-khai.md) | Checklist cho dev BackOffice TVBT |
| [nghiep-vu/01-tai-khoan-nhom-1.md](nghiep-vu/01-tai-khoan-nhom-1.md) | Mở/đóng TK ký quỹ; đăng ký/hủy TKGD tại TVKBT (MT598 600–603) |
| [nghiep-vu/02-ky-quy-nhom-2.md](nghiep-vu/02-ky-quy-nhom-2.md) | Nộp/rút ký quỹ bằng tiền & chứng khoán; phong tỏa; lãi tiền gửi |
| [nghiep-vu/03-bu-tru-vi-the-nhom-3.md](nghiep-vu/03-bu-tru-vi-the-nhom-3.md) | MT518; cảnh báo; chuyển khoản tất toán/vị thế; bù trừ omnibus |
| [nghiep-vu/04-thanh-toan-lai-lo-nhom-4.md](nghiep-vu/04-thanh-toan-lai-lo-nhom-4.md) | Nghĩa vụ VM và thanh toán T+1 |
| [nghiep-vu/05-thanh-toan-dao-han-nhom-5.md](nghiep-vu/05-thanh-toan-dao-han-nhom-5.md) | Đáo hạn chuyển giao vật chất và bằng tiền |
| [nghiep-vu/06-mat-kha-nang-thanh-toan-nhom-6.md](nghiep-vu/06-mat-kha-nang-thanh-toan-nhom-6.md) | Rút ký quỹ để thanh toán VM; thanh toán thay thế bằng tiền |
| [nghiep-vu/07-bao-cao-nhom-7.md](nghiep-vu/07-bao-cao-nhom-7.md) | Báo cáo tự động/theo yêu cầu; cấu trúc `.par` và 9 loại csv |

## Bản đồ nhanh: tài liệu gốc → file

| Tài liệu gốc | File |
|---|---|
| I. Phương thức trao đổi điện | 01 |
| I.1 Cấu trúc file FIN (1.1–1.4) | 02 |
| I.2 Ký tự hợp lệ & chuyển đổi tiếng Việt | 03 (+05) |
| II.1 Đăng ký TK giao dịch CKPS | nghiep-vu/01 |
| II.2 Nộp/rút ký quỹ (2.1–2.5) | nghiep-vu/02 |
| II.3 Bù trừ và quản lý vị thế (3.1–3.8) | nghiep-vu/03 |
| II.4 Thanh toán lãi/lỗ (4.1–4.2) | nghiep-vu/04 |
| II.5 Thanh toán đáo hạn (5.1–5.2) | nghiep-vu/05 |
| II.6 Xử lý mất khả năng thanh toán (6.1–6.2) | nghiep-vu/06 |
| II.7 Báo cáo (7.1–7.3) | nghiep-vu/07 |
| Bảng 1 / Bảng 2 | 04 / 05 |

## Ghi chú về tài liệu gốc

File Word chứa 23 sơ đồ luồng dạng ảnh vector (EMF/OLE) — không còn trong markdown dưới dạng ảnh; thông tin trong sơ đồ đã được đưa vào các sơ đồ Mermaid ở từng mục. Nếu cần đối chiếu hình thức, mở `Phu-luc-06.docx`.
