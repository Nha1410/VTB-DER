# Lưu ý và điểm chưa nhất quán trong văn bản gốc

> Các file khác trong `docs/` giữ **nguyên văn** tài liệu gốc (kể cả lỗi). File này liệt kê những chỗ cần làm rõ với VSD/đối tác hoặc cần thận trọng khi cài đặt. Mức độ: **Cao** = có thể ảnh hưởng tới việc gửi/nhận điện; **TB** = mâu thuẫn tài liệu cần xác nhận; **Thấp** = lỗi chính tả/trình bày.

## 1. Phần I — Phương thức, cấu trúc điện

| # | Mức | Vị trí | Vấn đề |
|---|---|---|---|
| 1 | TB | Phần I, ý 1 | Viện dẫn "Mục 3.1 / 3.2" nhưng nội dung nằm ở mục 1 (cấu trúc FIN) và mục 2 (ký tự/tiếng Việt) |
| 2 | Thấp | Tiêu đề Phần I | "ĐIỆN ĐIỆN NGHIỆP VỤ" (lặp từ) |
| 3 | Thấp | Thư mục `send` | "hệ thống GateWay của Thành viên bù trừ" — theo ngữ cảnh là GateWay của VSD |
| 4 | TB | 1.2 | Câu mở đầu ghi "điện FIN gửi lên hệ thống VSD" nhưng mục mô tả điện gửi **về** TVBT |
| 5 | TB | 1.1 vs 1.4 | Block 1 chuẩn: Application Identifier `F`. Ví dụ NAK dùng `{1:A21…}` (chữ `A`), ví dụ ACK dùng `{1:F21…}` |
| 6 | TB | 1.4 | Định dạng tag `177`: ví dụ NAK `9703051524` (yymmddhhmm), ví dụ ACK/NAK về TVBT `20140314 17:28:37`. Mô tả chỉ ghi "date-time" |
| 7 | TB | 1.4 | Ví dụ NAK dùng mã `H80` (không có trong Bảng 1); ví dụ khác dùng text `NAK [REQUESTID: duplicate]`. Bảng 1 có dòng `[content]` = nội dung lỗi cụ thể và `T98` = Duplicate message |
| 8 | Thấp | 1.4 | Ví dụ ACK/NAK có `}}}` thừa/ghép liền Block 1 của điện gốc; chép nguyên văn |
| 9 | TB | 2.2 | Nói ánh xạ theo "telex hoặc mã 4 ký tự unicode" nhưng Bảng 2 chỉ có quy tắc telex. Ví dụ VD1 viết `KH?OS?A`, VD2 viết `C?OO?NG` — chưa rõ khi nhiều ký tự có dấu liền nhau thì bọc `?` riêng từng ký tự hay gộp |
| 10 | Thấp | Bảng 2 | Ký tự `/` nằm trong bộ ký tự hợp lệ nhưng vẫn có mã `?_?` |

## 2. Nhóm 1–2 — Tài khoản, ký quỹ

| # | Mức | Vị trí | Vấn đề |
|---|---|---|---|
| 11 | TB | 2.4 | Văn bản: không thực hiện được → gửi **NAK**; sơ đồ: **MT548** từ chối. (2.3 ghi nhận cả hai: NAK khi sai cấu trúc/thiếu số dư, MT548 khi cán bộ VSD từ chối duyệt) |
| 12 | Thấp | 2.3 MT548 | Dòng cuối bảng ghi "Bắt đầu Block: Thông tin chung" (đáng lẽ "Kết thúc") |
| 13 | TB | 2.3 MT524 | Dòng 6 (`20C`, qualifier `4!c`, Field `4!c`) — qualifier không cụ thể; theo mô tả liên kết tới yêu cầu trước khi `23G`=`CANC` |
| 14 | TB | 2.3 MT508 | `36B` dùng qualifier `ESTT`, trong khi MT524 dùng `SETT` |
| 15 | Thấp | MT544/MT546 (2.4, 5.1.1) | Cột **No.** của MT544 bị lặp/lùi (…24, rồi 21, 22, 23…). MT542 ở 5.1.1 cũng lặp số (TRADDET đánh lại từ 6) |
| 16 | TB | MT546 (2.4) | `36B` qualifier `SETT` nhưng mô tả giá trị `:ESTT//UNIT`. MT546 (5.1.1) dùng `ESTT` |
| 17 | TB | 2.1 | Lỗi hạch toán: VSD gửi `MT198` cho VIETINBANK nhưng không có bảng đặc tả MT198 |
| 18 | Thấp | 2.1 | Lưu ý "NĐT A  NĐT B" thiếu ký hiệu mũi tên (mất khi chuyển Word) |
| 19 | TB | 2.1 MT910 | `72` quy định `IM//001C123456` — định dạng nội dung diễn giải bắt buộc để VSD nhận diện tài khoản |

## 3. Nhóm 3 — Bù trừ và vị thế

| # | Mức | Vị trí | Vấn đề |
|---|---|---|---|
| 20 | TB | 3.1 | Bảng MT598 `TRADE` dùng Sub-type 621; mô tả csv "Mã CKPS, Số TKGD ký quỹ, số hiệu TVBT, thời điểm khớp lệnh, giá khớp, số lượng" **khác** cấu trúc csv `TRADE` ở 7.3.2.1 (id, state, reference, settlement, time, account, buy/sell, code, instrument name, quantity, currency, price) |
| 21 | TB | 3.1 MT518 | `22H BUSE` mô tả "BUYI / SELL / CROF"; hai tag `95R SELL`/`BUYR` quy định VSD điền `VSDSVN06` cho phía không phải TVBT |
| 22 | TB | 3.1, 4.1, 7.3.1 | Bảng `.par`: `Requestor`/`Responder` ở 3.1 và 4.1 = (Requestor: VSD, Responder: TVBT) nhưng ở 7.3.1 mô tả ngược lại (Requestor = TVBT, Responder = VSD) dù cùng giá trị `o=<$VSDBICCODE>` / `o=<$MEMBERBICCODE>` |
| 23 | Thấp | `.par` | `Size`: "integer <= 0" (có thể là >= 0 hoặc > 0); `PossibleDuplicate` mặc định TRUE trong khi các cờ khác mặc định FALSE; `TransferRef` ở 7.3.1 chỉ ghi "POSITION" |
| 24 | Thấp | 3.2 | "tài sản lý quỹ" (lỗi gõ); khối `BCOL`/`SCOL` chú thích "Kết thúc Block: BCOL/SCOL" |
| 25 | TB | 3.2–3.5 | Cùng Sub-type 699 cho 4 loại cảnh báo, phân biệt bằng `77E` (COLLAT/POSITION/LIMIT/BLOCK) và `70D EXPO` (Y/X/S/1/2) |
| 26 | TB | 3.5 | Mục không có phần mô tả văn bản, chỉ có bảng |
| 27 | TB | 3.6–3.7 | Cùng bảng MT598 631 `GIVEUP`/`TAKEUP`; 3.7 không phát sinh MT546/544/900/910. Sơ đồ đặt tên "thông báo chuyển vị thế / nhận vị thế" |
| 28 | Thấp | 3.6 | Mũi tên "TVBT bên chuyển → …" bị mất khi chuyển Word (xuất hiện ở các bước 4.1, 4.2) |
| 29 | Thấp | 3.8 | "Bước 1; TVBT gửi điện M598" (gõ thiếu chữ T) |
| 30 | TB | 3.8 | Từ chối: sub-type **633** (không phải 631); lý do nằm ở `20C STAT`: "Vị thế không có" hoặc "Khối lượng vị thế nhỏ hơn khối lượng yêu cầu bù trừ" |

## 4. Nhóm 4–6 — Thanh toán

| # | Mức | Vị trí | Vấn đề |
|---|---|---|---|
| 31 | TB | 4.2 | MT910 "Thông báo chuyển tiền vào TK tiền thanh toán TVBT lãi" ghi chú "(có thể được thay đổi)" — nội dung chưa chốt. Các MT900/MT910 có chỗ `72` là `CRED`/`SETL`/`PAYMENT` với độ dài `35x`/`73x`/`6*35x` không đồng nhất |
| 32 | TB | 4.2 | Bước 1 có MT103 VSD→VIETINBANK (thông báo thực hiện thanh toán) nhưng không có bảng đặc tả |
| 33 | Thấp | 4.2 | Mũi tên giữa các tài khoản bị mất ("TK tiền thanh toán TVBT bên trả TK tiền thanh toán VSD") |
| 34 | Thấp | 4.1 | "Lấy giá trị621" (thiếu khoảng trắng) |
| 35 | TB | 5.1.2 | MT546/MT544 mục này ghi `22F STCO` với giá trị `:SETR//PHYS` (khả năng nên là `:STCO//PHYS`) |
| 36 | TB | 5.1.2 | Báo cáo PHYSICAL gửi cuối **E+1**; báo cáo BOND và MT598 nghĩa vụ thanh toán chiều **E+2**; thanh toán cuối **E+3**. Không có bảng MT598 "Thông báo nghĩa vụ thanh toán đáo hạn" riêng (giả định giống báo cáo BOND 621) |
| 37 | TB | 5.1.1 | Trường hợp bên bán mất khả năng: vay chứng khoán qua SBL; STCO `SPDL` trong MT546 — nêu trong mô tả nhưng không có bảng cho giao dịch SBL |
| 38 | Thấp | 5.1.2 | Bảng MT900/MT910 đầu mục ghi "Tương tự như…4.1/4.2" và mô tả ngầm theo mục cũ |
| 39 | TB | 6.1 | Văn bản bước 7 "MT900 thông báo rút tiền thành công", sơ đồ có thêm MT910 thông báo tăng tiền TK thanh toán. Bước 2 (công văn) là xử lý ngoài hệ thống |
| 40 | TB | 6.2 | Bước 1 (công văn VSD → NHTT, nội dung `TTDM`) là xử lý ngoài hệ thống. Bảng MT900/910 còn sót ghi chú nội bộ: "- (số tham chiếu từ hệ thống ngoài vào VTB) -> VTB check lại / - Trường hợp không có giá trị lấy F20" |
| 41 | Thấp | 6.2 | Câu "VSD gửi cho công văn đề nghị NHTT" thiếu tân ngữ |

## 5. Nhóm 7 — Báo cáo

| # | Mức | Vị trí | Vấn đề |
|---|---|---|---|
| 42 | **Cao** | 7.1/7.2 | Bảng "báo cáo theo yêu cầu" liệt kê **6 báo cáo** (TRADE, POSITION, COLLATERAL, MARGIN, SETTLEMENT, PHYSICAL) nhưng đoạn văn dưới đó nói tra cứu theo thời gian thực **loại trừ** SETTLEMENT và PHYSICAL (và bảng MT598 620 cũng loại BOND/INST_TRADE/INST_COLLAT). Cần xác nhận với VSD |
| 43 | TB | 7.2 | Không có đề mục "7.2" (chỉ là một mục danh sách); "7.3 Cấu trúc báo cáo" tiếp ngay sau |
| 44 | TB | 7.1 | Hai tiêu đề "Báo cáo từ số 1 và số 7" / "số 8 và số 9": báo cáo 1–7 dùng MT598 621; 8–9 dùng MT598 605. Mã báo cáo số 7 viết "BOND" cả ở cột tên |
| 45 | TB | 7.1 | Báo cáo 8 và 9 đều có tên "Confirmation Message Report" (khác mã INST_TRADE / INST_COLLAT) |
| 46 | TB | 7.3.2 | Cột Content của nhiều bảng csv trống; không nêu dấu phân cách, header, encoding, định dạng số/ngày (trừ `YYYYMMDD` ở INST_*) |
| 47 | Thấp | 7.3.2 | "Harcut", "Mã kỹ quỹ", "Số lượng kỹ quỹ" |
| 48 | TB | 7.1 | Tên file báo cáo `[Date]_[Report name]_[Recipient]` (ví dụ `20160323_Trade_001`); với INST_*: `[Date]_INST_TRADE_CCPIG_[Recipient]`. Tại MT598 621 báo cáo khác ví dụ `20160323_Reportcode_001` / `20160323_Settlement_001` — quy ước viết hoa/thường (Trade vs TRADE) không thống nhất |

## 6. Chính tả / trình bày (không ảnh hưởng nghiệp vụ)

- "kỹ quỹ", "Cá nhận trong nước", "thé vị", "chứng khoán phát sinh" (đúng: phái sinh), "BẢNG MÃ LÔI TRẢ VỂ" (đúng: lỗi trả về), "Cá nhận".
- Nhiều chỗ mất ký hiệu mũi tên "→" khi chuyển Word (ví dụ 2.1, 3.6, 4.2).
- Một số bảng thiếu số thứ tự trường liên tục (xem mục 15).
- Tên ngân hàng khi thì "NHTT"/"VIETINBANK", khi thì "VTB" (ghi chú nội bộ ở mục 6.2).

## 7. Điểm nên xác nhận với VSD trước khi cài đặt (gợi ý)

1. Định dạng chính xác tag `177` và cách phân biệt mã lỗi `405` (H/T/số/`[content]`).
2. Cách bọc `?` khi chuyển nhiều ký tự tiếng Việt liền nhau.
3. Mục 7.2: SETTLEMENT/PHYSICAL có tra cứu theo yêu cầu được không.
4. Tên file `.csv`/`.par` và quy ước tên (chữ hoa/thường), mã hóa, header csv.
5. Đặc tả MT598 "Thông báo nghĩa vụ thanh toán đáo hạn" (5.1.2), MT103 V→NHTT, MT198.
6. Giá trị `STCO` trong MT544/546 mục 5.1.2.

## 8. Phát hiện khi cài đặt bộ mã hoá (plan 001)

| # | Mức | Vị trí | Vấn đề |
|---|---|---|---|
| 49 | **Cao** | 2.1 vs Bảng 2 | Bảng 2 chuyển `/` thành `?_?`, `&` thành `?_38?`… tức dùng ký tự `_`, nhưng bộ ký tự hợp lệ ở 2.1 **không có `_`**. Code chấp nhận `_` chỉ bên trong cặp `?…?` |
| 50 | **Cao** | 2.1 vs 2.2 | `?` vừa là ký tự hợp lệ vừa là marker chuyển đổi; spec không nói cách gửi một dấu `?` thật. Code từ chối `?` thô khi mã hoá |
| 51 | TB | 1.1 Block 4 | Spec không nêu ký tự xuống dòng giữa các dòng của Block 4 (ví dụ chỉ có `…`). Code dùng CRLF theo chuẩn SWIFT FIN |
| 52 | TB | Bảng 2 | Spec không có mã cho `@` (email phải viết `(at)` theo bảng MT598) và `_` thô |
| 53 | Thấp | Bảng 2 | Bảng có 139 dòng dữ liệu (chữ cái tiếng Việt và 5 ký tự đặc biệt); số dòng này được cố định trong test để phát hiện khi tài liệu thay đổi |
