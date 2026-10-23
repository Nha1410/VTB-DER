# Checklist triển khai cho BackOffice (gợi ý từ tài liệu)

> **Đây là hướng dẫn tổng hợp, không phải nguyên văn Phụ lục 06.** Mỗi mục trỏ về nơi tài liệu quy định. Những chỗ tài liệu chưa nêu được đánh dấu **[Chưa quy định]** — xem [07-luu-y-va-diem-chua-nhat-quan.md](07-luu-y-va-diem-chua-nhat-quan.md).

## A. Kết nối và thư mục ([01](01-phuong-thuc-trao-doi.md))

- [ ] Cài GateWay Client; cấu hình thư mục nghiệp vụ **phái sinh** trong `\ext\com.fss.esb.client.properties`.
- [ ] Đăng nhập theo phiên bằng username/password + Token key (chữ ký số doanh nghiệp BKAV); GateWay Client tự thống nhất session key.
- [ ] BackOffice **ghi điện vào `send`**; **quét liên tục `receive`**; theo dõi `archive` (gửi thành công) và `error` (gửi lỗi).
- [ ] **[Chưa quy định]** Quy tắc đặt tên file `.fin`, cơ chế khóa file khi ghi/đọc, chu kỳ quét của GateWay Client.

## B. Tạo điện gửi lên VSD ([02](02-cau-truc-dien-fin.md))

- [ ] Block 1: `{1:F01<BICCODE TVBT><session 4 số><ISN 6 số>}`, ISN tăng dần, **không trùng trong 1 phiên**.
- [ ] Block 2: `{2:I<MT><VSDSVN06XXXX><N|U>…}`; nếu priority `U` thì delivery monitoring `1` hoặc `3`; nếu `N` thì `2` hoặc để trống.
- [ ] Block 4: đúng cấu trúc từng loại điện ([nghiep-vu/](nghiep-vu/)); Block 5: dùng mẫu `{5:{MAC:00000000}{CHK:F1DBCA886BBF}{TNG:}}`.
- [ ] BICCODE TVBT: `VSD` + 3 ký tự tên viết tắt + `X` đệm (ví dụ `VSDSSIXX`); BICCODE VSD/CCP: `VSDSVN06`.
- [ ] Số hiệu tham chiếu `:20:` (MT598/MT103) hoặc `:20C::SEME//` (MT5xx) phải **duy nhất** — VSD.STP kiểm soát trùng (NAK `T98` / `duplicate`).

## C. Dữ liệu và ký tự ([03](03-ky-tu-va-chuyen-doi-tieng-viet.md))

- [ ] Chỉ dùng ký tự hợp lệ; chuyển đổi tiếng Việt theo Bảng 2 (`?AAS?`, `?dd?`…), `@` trong email → `(at)`.
- [ ] Kiểm tra độ dài trường **sau khi** chuyển đổi (ví dụ `4*35x`, `35x`).
- [ ] Ngày dạng `YYYYMMDD` (`:4!c//8!n`); thời điểm khớp `8!n6!n` (`98C`); ngày/tiền/số tiền trong MT900/910 dạng `6!n3!a15d` (yymmdd + loại tiền + số tiền, dấu thập phân là dấu phẩy theo chuẩn SWIFT — **[Chưa quy định rõ trong tài liệu]**).
- [ ] Bên nhận chuyển đổi ngược `?XX?` → tiếng Việt khi hiển thị/lưu.

## D. Xử lý phản hồi

- [ ] Nhận **ACK/NAK** (Block 1 service `21`; tag `451`: 0 = ACK, 1 = NAK; tag `405` = mã lỗi): đối chiếu với điện đã gửi, đối chiếu mã lỗi theo [Bảng 1](04-bang-ma-loi.md).
- [ ] Liên kết điện phản hồi với yêu cầu qua các tham chiếu: `20C::RELA//` (MT5xx/MT598 phản hồi), `21` (MT900/MT910 trỏ `20` của MT103), `20C::PREV//` (hủy/điện trước).
- [ ] Trạng thái nghiệp vụ trong MT598 phản hồi: `25D::IPRC//PACK` (đồng ý) / `REJT` (từ chối), lý do `70D::REAS`. MT548: `25D::IPRC//REJT`, `24B::REJT/NARR`.
- [ ] Phân biệt từ chối **mức hệ thống** (NAK: sai cấu trúc, thiếu số dư, trùng) và **mức nghiệp vụ** (MT548/MT598 từ chối do cán bộ VSD).
- [ ] Hủy yêu cầu: MT524 `23G=CANC` / MT508 `23G=CANC` (liên kết `LINK//524` hoặc `508`).

## E. Điện chủ động từ VSD cần xử lý (không có yêu cầu trước)

| Điện | Hành động đề xuất |
|---|---|
| MT518 (tức thời) | Ghi nhận giao dịch thế vị; `23G=CANC` thì hủy MT518 cũ (`20C::PREV`) |
| MT598 699 cảnh báo | Xử lý theo `70D::EXPO` (Y/X/S/1/2) — hiển thị cảnh báo, chặn giao dịch tài khoản/mã nếu `X`/`S` |
| MT598 621 `TRADE`/`SETTLEMENT`, 605, 621 báo cáo | Đọc tên csv ở `20C::PREV`; tải `.csv`/`.par` đính kèm theo [nhóm 7](nghiep-vu/07-bao-cao-nhom-7.md) |
| MT598 631 `GIVEUP`/`TAKEUP` | Cập nhật vị thế (long/short, giá trung bình) |
| MT910/MT900 từ VSD | Cập nhật số dư ký quỹ/thanh toán theo `72` |
| MT544/MT546 | Cập nhật số dư chứng khoán ký quỹ |

## F. Lịch xử lý định kỳ (tổng hợp)

| Thời điểm | Sự kiện | Tài liệu |
|---|---|---|
| Trong ngày, tức thời | MT518; cảnh báo MT598 699 | 3.1–3.5 |
| Cuối ngày T | MT598 `TRADE`, `SETTLEMENT` + csv; báo cáo POSITION, COLLATERAL, MARGIN; INST_TRADE, INST_COLLAT | 3.1, 4.1, 7.1 |
| T+1 | NHTT thực hiện chuyển khoản thanh toán VM; MT900/MT910 | 4.2 |
| Cuối E+1 | Báo cáo PHYSICAL | 7.1 |
| Chiều E+2 | MT598 nghĩa vụ thanh toán đáo hạn + báo cáo BOND | 5.1.2, 7.1 |
| E+3 | Chuyển giao chứng khoán, thanh toán tiền (MT546/544, MT900/910) | 5.1.2 |
| Cuối tháng | NHTT tính và hạch toán lãi tiền gửi ký quỹ; ngày liền sau chuyển về TK `BICCODE.R` | 2.5 |

## G. Kiểm thử đề xuất

- [ ] Gửi từng loại điện yêu cầu (600, 602, MT103 rút ký quỹ, MT524, MT542, 630, 620) với dữ liệu hợp lệ và không hợp lệ (trùng tham chiếu, thiếu trường `M`, sai độ dài, ký tự cấm).
- [ ] Mô phỏng nhận từng loại điện phản hồi/chủ động (xem [06](06-danh-muc-dien-mt.md)).
- [ ] Kiểm tra chuyển đổi tiếng Việt hai chiều bằng toàn bộ Bảng 2 và hai ví dụ VD1/VD2.
- [ ] Kiểm tra phân tích file `.par` và `.csv` các báo cáo ([nhóm 7](nghiep-vu/07-bao-cao-nhom-7.md)).
