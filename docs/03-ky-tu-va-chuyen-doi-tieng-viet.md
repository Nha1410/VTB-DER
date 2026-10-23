# Bộ ký tự hợp lệ và quy tắc chuyển đổi tiếng Việt

> Nguồn: Phụ lục 06 — Phần I, mục 2 (2.1–2.2) và các ví dụ ở cuối tài liệu. Bảng tra cứu đầy đủ: [05-bang-chuyen-doi-ky-tu.md](05-bang-chuyen-doi-ky-tu.md).

## 2.1 Bộ ký tự hợp lệ

Các ký tự sau được coi là hợp lệ trong nội dung điện nghiệp vụ:

| Nhóm | Ký tự |
|---|---|
| Chữ thường | `a b c d e f g h i j k l m n o p q r s t u v w x y z` |
| Chữ hoa | `A B C D E F G H I J K L M N O P Q R S T U V W X Y Z` |
| Chữ số | `0 1 2 3 4 5 6 7 8 9` |
| Ký tự đặc biệt | `/ - ? : ( ) . , ' +` |
| Điều khiển | `<Cr>` (carriage return), `<Space>` |

Mọi ký tự ngoài tập này (kể cả chữ tiếng Việt có dấu, `&`, `#`, `%`, `\`, `@`…) **không được xuất hiện trực tiếp** và phải chuyển đổi theo mục 2.2. Riêng ký tự `/` cũng có mã chuyển đổi (`?_?`) trong Bảng 2 dù nằm trong danh sách hợp lệ.

## 2.2 Quy tắc chuyển đổi tiếng Việt

- Ký tự tiếng Việt **không được phép xuất hiện trực tiếp**, phải chuyển thành ký tự latin theo bảng quy đổi (Bảng 2).
- **Bên gửi** chuyển đổi tiếng Việt → latin; **bên nhận** chuyển đổi ngược lại.
- Để báo cho bên nhận biết cần chuyển đổi, bên gửi thêm ký tự **chấm hỏi `?` vào trước và sau** ký tự cần chuyển đổi.
- Mỗi ký tự tiếng Việt cần chuyển đổi ánh xạ **một–một** với nhóm ký tự latin tương ứng (theo nguyên tắc đánh máy **telex** của bộ gõ tiếng Việt đang áp dụng, hoặc mã 4 ký tự unicode Basic Latin và Latin-1 Supplement — tham khảo https://en.wikipedia.org/wiki/List_of_Unicode_characters).
- **Độ dài ký tự được tính trên chuỗi latin sau khi chuyển đổi** theo Bảng 2 (quan trọng khi kiểm tra giới hạn độ dài trường như `16x`, `35x`, `4*35x`…).

### Quy luật rút ra từ Bảng 2 (diễn giải, không phải nguyên văn)

| Thành phần | Quy tắc telex |
|---|---|
| Chữ cái biến thể | `ă`=`aw`, `â`=`aa`, `ê`=`ee`, `ô`=`oo`, `ơ`=`ow`, `ư`=`uw`, `đ`=`dd` |
| Dấu thanh | huyền=`f`, sắc=`s`, hỏi=`r`, ngã=`x`, nặng=`j` |
| Cách ghép | chữ cái (biến thể) + dấu thanh, bọc trong `?…?`; hoa/thường giữ nguyên chữ hoa/thường (`?AAS?` = Ấ, `?aas?` = ấ) |
| Ký tự đặc biệt | `/` → `?_?`; `&` → `?_38?`; `#` → `?_35?`; `%` → `?_37?`; `\` → `?_92?` (số là mã ASCII thập phân) |

### Ví dụ (nguyên văn)

**VD1**

| Bước | Chuỗi | Độ dài |
|---|---|---|
| Chuỗi gốc bên gửi tạo trước khi chuyển đổi | `KHÓA` | 4 ký tự |
| Điện ISO 15022 sau khi chuyển đổi | `KH?OS?A` | 7 ký tự |
| Bên nhận chuyển đổi lại | `KHÓA` | 4 ký tự |

**VD2**

| Bước | Chuỗi |
|---|---|
| Chuỗi gốc bên gửi tạo | `CÔNG TY SỮA` |
| Điện ISO 15022 sau khi chuyển đổi | `C?OO?NG TY S?UWX?A` |
| Bên nhận chuyển đổi lại | `CÔNG TY SỮA` |

> Lưu ý quan trọng khi cài đặt: ở VD1 chữ `Ó` được viết `?OS?` (bọc `?` quanh **từng ký tự** cần chuyển đổi), nhưng ở VD2 `Ô` → `?OO?` và `Ữ` → `?UWX?`. Văn bản gốc không có ví dụ cho hai ký tự có dấu đứng liền nhau; khi gặp, nên hỏi/kiểm tra với VSD xem có gộp `?` hay không.
