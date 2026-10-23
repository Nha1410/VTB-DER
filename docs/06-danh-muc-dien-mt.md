# Danh mục điện MT — tra cứu nhanh

> Tổng hợp từ Phụ lục 06. Mỗi dòng trỏ tới mục có bảng đặc tả chi tiết. Chiều: **T→V** = TVBT gửi VSD; **V→T** = VSD gửi TVBT; **V→N** = VSD gửi NHTT; **N→V** = NHTT gửi VSD. "(không có bảng)" = điện chỉ được nhắc trong trình tự/sơ đồ, tài liệu không có bảng đặc tả.

## 1. Điện MT598 (điện tự do) — phân loại theo Sub-Message Type (tag `12`) và `77E`

| Sub-type (tag 12) | `77E` Proprietary | Chiều | Mục đích | Mục |
|---|---|---|---|---|
| 600 | `/TAM` | T→V | Yêu cầu mở/đóng TK giao dịch cho ký quỹ & giao dịch CKPS | [1.1](nghiep-vu/01-tai-khoan-nhom-1.md) |
| 601 | `/TAM` | V→T | Xác nhận kết quả mở/đóng TK | 1.1 |
| 602 | `/TAT` | T→V | Yêu cầu đăng ký/hủy đăng ký TKGD tại TVKBT | 1.2 |
| 603 | `/TAT` | V→T | Xác nhận kết quả đăng ký/hủy đăng ký | 1.2 |
| 613 | `/CASH` | V→T | Từ chối yêu cầu rút tiền ký quỹ | [2.2](nghiep-vu/02-ky-quy-nhom-2.md) |
| 621 | `TRADE` | V→T | Tổng hợp kết quả giao dịch thế vị cuối ngày (+ file .csv/.par) | [3.1](nghiep-vu/03-bu-tru-vi-the-nhom-3.md) |
| 699 | `COLLAT` | V→T | Cảnh báo tỷ lệ sử dụng tài khoản ký quỹ | 3.2 |
| 699 | `POSITION` | V→T | Cảnh báo vi phạm giới hạn vị thế | 3.3 |
| 699 | `LIMIT` | V→T | Cảnh báo vi phạm giới hạn khối lượng mở | 3.4 |
| 699 | `BLOCK` | V→T | Cảnh báo phong tỏa tài khoản | 3.5 |
| 631 | `GIVEUP` / `TAKEUP` | V→T | Thông báo giảm/tăng vị thế (chuyển khoản tất toán, chuyển khoản vị thế) | 3.6, 3.7 |
| 630 | `REDUCE` | T→V | Yêu cầu bù trừ vị thế (omnibus) | 3.8 |
| 631 | `REDUCE` | V→T | Xác nhận yêu cầu bù trừ vị thế | 3.8 |
| 633 | `REDUCE` | V→T | Từ chối yêu cầu bù trừ vị thế | 3.8 |
| 621 | `SETTLEMENT` | V→T | Thông báo nghĩa vụ thanh toán lãi/lỗ VM (+ file .csv/.par) | [4.1](nghiep-vu/04-thanh-toan-lai-lo-nhom-4.md) |
| — | (không có bảng) | V→T | Thông báo nghĩa vụ thanh toán đáo hạn kèm file csv (chiều E+2) | [5.1.2](nghiep-vu/05-thanh-toan-dao-han-nhom-5.md) |
| 605 | `INST_TRADE` / `INST_COLLAT` | V→T | Báo cáo danh mục chứng khoán (được phép giao dịch / nhận làm TSĐB) | [7.1](nghiep-vu/07-bao-cao-nhom-7.md) |
| 621 | mã báo cáo (`TRADE`, `POSITION`, `COLLATERAL`, `MARGIN`, `SETTLEMENT`, `PHYSICAL`, `BOND`) | V→T | Thông tin chi tiết báo cáo (+ csv); cũng là phản hồi tra cứu | 7.1, 7.2 |
| 620 | tham số báo cáo (`20*78x`) | T→V | Yêu cầu tra cứu báo cáo theo yêu cầu | 7.2 |

## 2. Các điện MT khác

| MT | Tên (theo tài liệu) | Chiều | Mục |
|---|---|---|---|
| MT103 | Yêu cầu rút tiền ký quỹ | T→V | [2.2](nghiep-vu/02-ky-quy-nhom-2.md) |
| MT103 | Yêu cầu NHTT chuyển tiền rút ký quỹ (không có bảng) | V→N | 2.2 |
| MT103 | Thông báo thực hiện thanh toán VM cho từng TVBT (không có bảng) | V→N | [4.2](nghiep-vu/04-thanh-toan-lai-lo-nhom-4.md) |
| MT103 | Yêu cầu rút tiền ký quỹ để thanh toán VM | T→V (bảng) / V→N (trình tự bước 4) | [6.1](nghiep-vu/06-mat-kha-nang-thanh-toan-nhom-6.md) |
| MT198 | Phản hồi lỗi hạch toán nộp ký quỹ — kèm mã lỗi (không có bảng) | V→N | 2.1 |
| MT508 | Xác nhận phong tỏa/giải tỏa chứng khoán ký quỹ | V→T | 2.3 |
| MT518 | Thông báo kết quả giao dịch đã được thế vị | V→T | 3.1 |
| MT524 | Yêu cầu phong tỏa/giải tỏa chứng khoán ký quỹ | T→V | 2.3 |
| MT542 | Yêu cầu nộp/rút TSKQ bằng chứng khoán (2.4); yêu cầu nộp chứng khoán chuyển giao vật chất (5.1.1) | T→V | 2.4, 5.1.1 |
| MT544 | Thông báo tăng số dư / xác nhận nhận chứng khoán | V→T | 2.4, 3.6, 5.1.1, 5.1.2 |
| MT546 | Thông báo giảm số dư / xác nhận chuyển chứng khoán | V→T | 2.4, 3.6, 5.1.1, 5.1.2 |
| MT548 | Từ chối yêu cầu phong tỏa/giải tỏa (2.3) hoặc nộp CK chuyển giao (5.1.1) | V→T | 2.3, 5.1.1 (2.4: sơ đồ ghi MT548, văn bản ghi NAK) |
| MT900 | Báo NỢ — rút tiền thành công / ghi giảm / cắt tiền / báo nợ chuyển giao | N→V, V→T | 2.2, 3.6, 4.2, 5.1.2, 6.1, 6.2 |
| MT910 | Báo CÓ — nộp ký quỹ / ghi tăng / chuyển tiền vào TK thanh toán | N→V, V→T | 2.1, 3.6, 4.2, 5.1.2, 6.1, 6.2 |

## 3. Tra cứu theo nghiệp vụ

| Nghiệp vụ | Điện chính |
|---|---|
| Mở/đóng TK giao dịch ký quỹ | MT598 600 → 601 |
| Đăng ký/hủy đăng ký TKGD tại TVKBT | MT598 602 → 603 |
| Nộp ký quỹ bằng tiền | NĐT/TVBT chuyển tiền vào NHTT → MT910 (NHTT→VSD) → MT910 (VSD→TVBT); lỗi: MT198 |
| Rút ký quỹ bằng tiền | MT103 → (MT103 V→N) → MT900 (N→V) → MT900 (V→T); từ chối: MT598 613 |
| Phong tỏa/giải tỏa CK ký quỹ | MT524 → MT508 / MT548 |
| Nộp/rút ký quỹ bằng CK | MT542 → MT546 + MT544 / NAK |
| Kết quả giao dịch | MT518 (tức thời) + MT598 621 `TRADE` (cuối ngày) |
| Cảnh báo | MT598 699 (`COLLAT`/`POSITION`/`LIMIT`/`BLOCK`) |
| Đổi TVBT | MT546+MT900+MT598 631 GIVEUP → TVBT chuyển; MT544+MT910+MT598 631 TAKEUP → TVBT nhận |
| Chuyển khoản vị thế | MT598 631 GIVEUP/TAKEUP |
| Bù trừ vị thế omnibus | MT598 630 → 631 / 633 |
| Thanh toán VM | MT598 621 `SETTLEMENT` + MT103 (V→N) → MT900/MT910 |
| Đáo hạn chuyển giao vật chất | MT542 → MT546/MT544/MT548; E+2 MT598; E+3 MT546/MT544 + MT900/MT910 |
| Mất khả năng thanh toán | 6.1: MT103 → MT900 → MT900/MT910; 6.2: MT900/MT910 (TTDM) |
| Báo cáo | tự động MT598 621/605; theo yêu cầu MT598 620 → 621 |

> Cột "Chiều" của MT103/MT900/MT910 được **suy luận** từ trình tự; xem bảng chi tiết trong từng mục để biết bên gửi/nhận chính xác.
