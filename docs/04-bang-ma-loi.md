# Bảng 1 — Mã lỗi trả về TVBT

> Nguồn: Phụ lục 06 — BẢNG 1.

Mã lỗi nằm trong tag `405` (rejection-reason) của điện NAK khi tag `451` = 1 (xem [02-cau-truc-dien-fin.md](02-cau-truc-dien-fin.md)).

| Thông tin lỗi | Chú giải |
|---|---|
| -1 | Không thể load được dữ liệu đã ký |
| -2 | Không tìm thấy chữ ký số trong dữ liệu |
| -3 | Chữ ký đã bị thay đổi |
| -4 | Đường dẫn tin tưởng(Trustpath) không hợp lệ |
| -5 | Không tìm thấy chứng thư số của CA |
| 1 | Chứng thư số chưa đến hạn sử dụng |
| 2 | Chứng thư số đã hết hạn |
| 3 | Chứng thư số không có quyền ký |
| 4 | Chứng thư số đã bị thu hồi |
| 5 | Không thể kiểm tra trạng thái chứng thư số |
| 6 | Không thể kiểm tra chứng thư số qua OCSP: Không tồn tại url trong chứng thư số |
| 7 | Không tìm thấy chứng thư số trong chữ ký |
| 8 | Không thể kiểm tra chứng thư số qua OCSP: response null |
| 9 | Không thể kiểm tra chứng thư số qua CRL |
| 10 | Không thể kiểm tra chứng thư số qua OCSP: url không tồn tại |
| 11 | Chữ ký trong OCSP response không hợp lệ |
| H01 | Header missing/block 1 error. |
| H02 | Header application invalid. |
| H25 | Header formatting error. |
| H30 | Header message type invalid. |
| H50 | Header invalid destination (receiver). |
| H98 | Header invalid source (sender). |
| H99 | Header error. |
| T02 | Unknown error (please contact support). |
| T27 | BIC code formatting error. |
| T28 | BIC code is invalid destination. |
| T31 | Line/Sub-field/Component missing or invalid. |
| T32 | Line/Sub-field/Component not found. |
| T33 | Line/Sub-field/Component too long. |
| T34 | Line/Sub-field/Component too short. |
| T38 | Invalid time. |
| T40 | Invalid amount/number. |
| T45 | Invalid BIC code. |
| T52 | Invalid currency code. |
| T98 | Duplicate message detected. |
| [content] | Nội dung lỗi cụ thể , ví dụ: Tài khoản không tồn tại |


