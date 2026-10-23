# Cấu trúc điện FIN

> Nguồn: Phụ lục 06 — Phần I, mục 1 ("Cấu trúc File FIN", 1.1–1.4). Chuyển thể nguyên văn, chỉ định dạng lại.

Các điện nghiệp vụ trao đổi giữa VSD và TVBT được thể hiện dưới dạng **điện FIN** (file dữ liệu `.fin`). Một điện FIN gửi lên VSD gồm các BLOCK:

```
{1: BASIC HEADER BLOCK}
{2: APPLICATION HEADER BLOCK}
{4: TEXT BLOCK}
{5: TRAILER BLOCK}
```

## 1.1 Điện FIN chuẩn gửi lên VSD (TVBT → VSD)

### Block 1 — `{1: BASIC HEADER BLOCK}`

Mẫu: `{1:F01BANKBEBBAXXX2222123456}` tách theo thành phần:

| (a) | (b) | (c) | (d) | (e) | (f) |
|---|---|---|---|---|---|
| `{1:` | `F` | `01` | `BANKBEBBAXXX` | `2222` | `123456}` |

| Thành phần | Tên | Giá trị |
|---|---|---|
| (a) | Block Identifier | Điền `1:` |
| (b) | Application Identifier | Điền `F` |
| (c) | Service Identifier | Điền `01` |
| (d) | Sender's Financial Entity Identifier with Logical Terminal Identifier (always A) | Giá trị **BICCODE của bên gửi** |
| (e) | Session Number | Mã phiên |
| (f) | Input Sequence Number (ISN) | Số thứ tự của điện gửi trong phiên (tự tăng) |

**Ghi chú:** Trong 1 phiên, TVBT phải đảm bảo các điện gửi lên **khác nhau về số thứ tự** trong phiên.

### Block 2 — `{2: APPLICATION HEADER BLOCK}`

Mẫu: `{2:I103VSDSVN06XXXXXXX}` tách theo thành phần:

| (a) | (b) | (c) | (d) | (e) | (f) | (g) |
|---|---|---|---|---|---|---|
| `{2:` | `I` | `103` | `VSDSVN06XXXX` | `X` | `X` | `X}` |

| Thành phần | Tên | Giá trị |
|---|---|---|
| (a) | Block Identifier | Điền `2:` |
| (b) | Input Identifier | Điền `I` |
| (c) | Message Type | Điền mã điện MT |
| (d) | Receiver's Financial Entity Identifier with Logical Terminal Identifier (always X) | Địa chỉ bên nhận, thông thường là **VSD BICODE** |
| (e) | Message Priority | `U` (Khẩn) hoặc `N` (Thông thường) |
| (f) | Delivery monitoring | `1` hoặc `3` cho Message Priority = `U`; `2` hoặc `<null>` cho Message Priority = `N` |
| (g) | Obsolescence Period (optionally) | Tùy chọn của Block 2. Nếu có giá trị sẽ điền là `020` |

### Block 4 — `{4: TEXT BLOCK}`

Điền nội dung điện. **Ghi chú:** tất cả các ví dụ nêu trong quy trình nghiệp vụ ([nghiep-vu/](nghiep-vu/)) chỉ đưa thông tin trong Block này.

### Block 5 — `{5: TRAILER BLOCK}`

Không quy định thông tin (điền theo mẫu trong ví dụ dưới).

### Ví dụ

```
{1:F01VSDCTBVNAXXX0020000021}{2:I500VSDSVN06XXXXN}{4:
…
-}{5:{MAC:00000000}{CHK:F1DBCA886BBF}{TNG:}}
```

## 1.2 Điện FIN gửi về cho TVBT (VSD → TVBT)

Gồm các BLOCK tương tự mục 1.1 nhưng khác ở Block 1 và Block 2. *(Văn bản gốc ghi "Một điện FIN gửi lên hệ thống VSD gồm…" — theo ngữ cảnh là điện gửi về TVBT.)*

### Block 1

Mẫu: `{1:F01VSDSVN06AXXX2222123456}`

| (a) | (b) | (c) | (d) | (e) | (f) |
|---|---|---|---|---|---|
| `{1:` | `F` | `01` | `VSDSVN06AXXX` | `2222` | `123456}` |

| Thành phần | Tên | Giá trị |
|---|---|---|
| (a) | Block Identifier | `1:` |
| (b) | Application Identifier | `F` |
| (c) | Service Identifier | `01` |
| (d) | Sender's Financial Entity Identifier with Logical Terminal Identifier (always A) | BICCODE của VSD là `VSDSVN06` |
| (e) | Output Session Number | 4 ký tự số đại diện cho phiên trao đổi |
| (f) | Output Sequence Number (OSN) | Số thứ tự của điện gửi trong phiên (tự tăng). **Trong 1 phiên, số thứ tự này là duy nhất.** |

### Block 2

Mẫu (các thành phần theo bảng gốc):

| (a) | (b) | (c) | (d) | (e) | (f) | (g) |
|---|---|---|---|---|---|---|
| `{2:` | `O` | `103` | `1511010606` | `BANKVNVVAXXX0325013085` | `010515` | `1149}` |

| Thành phần | Ý nghĩa (theo văn bản gốc) |
|---|---|
| (a) | Số hiệu block, lấy giá trị là `2` |
| (b) | `O` cho "output", hoặc nhận từ bên gửi |
| (c) | Loại điện |
| (d) | Thời điểm gửi tính theo mốc thời gian của bên gửi |
| (e) | INPUT SEQUENCE NUMBER: gồm BICCODE bên gửi, số hiệu phiên và số thứ tự của điện |
| (f) | Ngày nhận theo mốc thời gian bên nhận |
| (g) | Giờ nhận theo mốc thời gian bên nhận |

## 1.3 Quy định mã VSD BICCODE

**Mã VSD BICCODE cấp cho TVBT** gồm 8 ký tự, cấp theo nguyên tắc:

```
VSD + [3 ký tự tên viết tắt của TVBT] + [XX]
```

- Ví dụ: `VSDSSIXX`, `VSDHSCXX`.
- Nếu tên viết tắt của TVBT dài hơn 3 ký tự, các ký tự từ thứ 4 trở đi **thay vào các ký tự `X` phía sau**. Ví dụ: `VSDBVSCX`, `VSDACBSX`, `VSDFPTSX`.
- **Mã VSD BICCODE của VSD cho CCP: `VSDSVN06`.**

## 1.4 Điện ACK/NAK từ VSD

Nội dung điện ACK/NAK gồm 2 phần chính: **phần ACK/NAK** và **nội dung điện MT nghiệp vụ Client gửi lên**. Phần ACK/NAK gồm Block 1 (Basic Header) và Block 4 (Text Block).

### Block 1

Tương tự điện nghiệp vụ gửi từ TVBT lên VSD nhưng (c) lấy giá trị là `21`:

| (a) | (b) | (c) | (d) | (e) | (f) |
|---|---|---|---|---|---|
| `{1:` | `F` | `21` | `BANKBEBBAXXX` | `2222` | `123456}` |

### Block 4

| Tag | Field | Description |
|---|---|---|
| `177` | date-time | Thời gian điện ACK/NAK phản hồi |
| `451` | accept-reject | `0` = accepted (ACK) / `1` = rejected (NAK) |
| `405` | rejection-reason | Chi tiết lỗi theo [Bảng 1](04-bang-ma-loi.md) nếu giá trị của thẻ `451` = 1 |

### Ví dụ điện NAK (nguyên văn)

```
{1:A21VNDZBET2AXXX0018000015}{4:{177:9703051524}{451:1}{405:H80}}
```

| Phần | Ý nghĩa |
|---|---|
| `{1:A21VNDZBET2AXXX0018000015}` | Basic Header |
| `{4:{177:9703051524}` | Text Block |
| `{451:1}` | The message is rejected... |
| `{405:H80}}` | ...because of delivery option error (H80). |

### Ví dụ điện ACK gửi về TVBT

```
{1:F21VSDCTBXXAXXX0020000021}{4:{177:20140314 17:28:37}{451:0}}}{1:F01VSDCTBXXAXXX0020000021}{2:I598VSDSVN06AXXXN}{4:
…
-}{5:{MAC:00000000}{CHK:F1DBCA886BBF}{TNG:}}
```

### Ví dụ điện NAK gửi về TVBT

```
{1:F21VSDCTBXXAXXX0020000021}{4:{177:20140314 17:27:49}{451:1}{405:NAK
[REQUESTID: duplicate]}}}{1:F01VSDCTBXXAXXX0020000021}{2:I598VSDSVN06AXXXN}{4:
…
-}{5:{MAC:00000000}{CHK:F1DBCA886BBF}{TNG:}}
```

> Các ví dụ ở trên được chép nguyên văn (kể cả những điểm không nhất quán, ví dụ định dạng tag `177` khác nhau giữa các ví dụ). Xem [07-luu-y-va-diem-chua-nhat-quan.md](07-luu-y-va-diem-chua-nhat-quan.md).
