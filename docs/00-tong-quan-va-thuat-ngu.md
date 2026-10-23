# Tổng quan và thuật ngữ

> Nguồn: Phụ lục 06 (QĐ 115/QĐ-VSD ngày 31/05/2017). Mục "Diễn giải" đánh dấu phần **suy luận từ ngữ cảnh**, không phải nguyên văn tài liệu.

## 1. Phạm vi tài liệu

Phụ lục 06 quy định **phương thức trao đổi** và **chuẩn điện nghiệp vụ** (ISO 15022, dạng điện FIN/MT) giữa hệ thống BackOffice của **Thành viên bù trừ (TVBT)** và **Cổng giao tiếp (GateWay) của VSD**, cho hoạt động **bù trừ, thanh toán giao dịch chứng khoán phái sinh**.

| Phần | Nội dung | File |
|---|---|---|
| I | Phương thức trao đổi; cấu trúc điện FIN; ký tự hợp lệ; chuyển đổi tiếng Việt | [01](01-phuong-thuc-trao-doi.md), [02](02-cau-truc-dien-fin.md), [03](03-ky-tu-va-chuyen-doi-tieng-viet.md) |
| II.1 | Đăng ký thông tin tài khoản giao dịch CKPS cho NĐT | [nghiep-vu/01](nghiep-vu/01-tai-khoan-nhom-1.md) |
| II.2 | Nộp/rút ký quỹ | [nghiep-vu/02](nghiep-vu/02-ky-quy-nhom-2.md) |
| II.3 | Bù trừ và quản lý vị thế | [nghiep-vu/03](nghiep-vu/03-bu-tru-vi-the-nhom-3.md) |
| II.4 | Thanh toán lãi/lỗ vị thế | [nghiep-vu/04](nghiep-vu/04-thanh-toan-lai-lo-nhom-4.md) |
| II.5 | Thanh toán đáo hạn | [nghiep-vu/05](nghiep-vu/05-thanh-toan-dao-han-nhom-5.md) |
| II.6 | Xử lý mất khả năng thanh toán | [nghiep-vu/06](nghiep-vu/06-mat-kha-nang-thanh-toan-nhom-6.md) |
| II.7 | Báo cáo | [nghiep-vu/07](nghiep-vu/07-bao-cao-nhom-7.md) |
| Bảng 1, 2 | Mã lỗi; bảng chuyển đổi ký tự | [04](04-bang-ma-loi.md), [05](05-bang-chuyen-doi-ky-tu.md) |

## 2. Các bên tham gia

| Bên | Vai trò (theo tài liệu) |
|---|---|
| **TVBT** — Thành viên bù trừ | Gửi/nhận điện qua GateWay Client; có BICCODE dạng `VSD`+3 ký tự tên viết tắt+`XX` |
| **TVKBT / NCM** — Thành viên không bù trừ | Nơi NĐT có thể mở tài khoản giao dịch (TKGD), nhưng bù trừ trên tài khoản ký quỹ mở tại TVBT (mục 1.2). Sơ đồ ghi "NCM" |
| **VSD** — Trung tâm Lưu ký Chứng khoán Việt Nam | Đóng vai trò CCP; BICCODE `VSDSVN06`; vận hành hệ thống VSD.STP |
| **CCP** — Đối tác bù trừ trung tâm | "Mã VSD BICCODE của VSD cho CCP: VSDSVN06"; hệ thống CCP gửi báo cáo cuối ngày |
| **NHTT** — Ngân hàng thanh toán | VIETINBANK (VietinBank). Giữ tài khoản tiền ký quỹ/thanh toán đứng tên VSD |
| **NĐT** — Nhà đầu tư | Chủ tài khoản giao dịch/ký quỹ |
| **Sở giao dịch (HNX)** | Gửi kết quả khớp lệnh cho VSD để thế vị (mục 3.1). Báo cáo dùng "HNX trading code" |
| **SBL** | Hệ thống vay chứng khoán, dùng khi bên bán mất khả năng chuyển giao trái phiếu (mục 5.1.1) |

## 3. Thuật ngữ và từ viết tắt

| Từ | Nghĩa | Ghi chú |
|---|---|---|
| CKPS | Chứng khoán phái sinh | Tài liệu đôi chỗ viết "chứng khoán phát sinh" |
| HĐTL | Hợp đồng tương lai | Hợp đồng tương lai chỉ số cổ phiếu: đáo hạn bằng tiền. Hợp đồng tương lai trái phiếu chính phủ: bằng tiền **hoặc** chuyển giao vật chất (mục 5.2) |
| TPCP | Trái phiếu Chính phủ | |
| TKGD | Tài khoản giao dịch | |
| TKGDPS | Tài khoản giao dịch phái sinh | Dùng trong sơ đồ mục 2.3 |
| TSKQ | Tài sản ký quỹ | Diễn giải: dùng trong sơ đồ mục 3.6 |
| VM | Variation Margin — lãi/lỗ vị thế hàng ngày | Tài liệu: "Thanh toán lãi/lỗ vị thế (VM)" |
| IM | Initial Margin — ký quỹ ban đầu | Diễn giải; xuất hiện: `IM//số TK` (mục 2.1/2.2), `SETR//COLI/COLO` "ký quỹ ban đầu" |
| DSP | Daily Settlement Price — giá thanh toán hàng ngày | Diễn giải; xuất hiện: báo cáo POSITION cột "Value (DSP)", INST_TRADE cột "Daily Settlement Price" |
| OI | Open Interest — tổng khối lượng mở của thị trường | Mục 3.4 |
| E | Ngày đáo hạn hợp đồng | Chuyển giao vật chất: báo cáo PHYSICAL cuối E+1; báo cáo BOND và MT598 nghĩa vụ chiều E+2; thanh toán cuối cùng E+3 |
| T / T+1 | Ngày giao dịch / ngày kế tiếp | Mục 4.2: thông báo cuối ngày T, thanh toán T+1 |
| TVMKNTT | TVBT mất khả năng thanh toán | Mục 6.2 |
| STP | Straight-Through Processing | "Hệ thống VSD.STP kiểm soát tính loại điện và số hiệu tham chiếu phải là duy nhất" |
| BIC / BICCODE | Mã định danh ngân hàng/tổ chức (SWIFT) | Xem [02](02-cau-truc-dien-fin.md) mục 1.3 |
| MT | Message Type — loại điện ISO 15022/SWIFT | Xem [06](06-danh-muc-dien-mt.md) |
| ISIN | Mã quốc tế định danh chứng khoán | Định dạng `[ISIN1!e12!c]` trong tag `35B` |
| Omnibus (NETT//OMNI) | Tài khoản tổng hợp (not-netted) | Vị thế mua/bán cùng sản phẩm, cùng tháng đáo hạn vẫn giữ riêng; có thể yêu cầu bù trừ (mục 3.8) |
| Haircut | Tỷ lệ chiết khấu tài sản đảm bảo | Báo cáo COLLATERAL, INST_COLLAT |
| Give-up / Take-up | Ghi giảm / ghi tăng vị thế khi chuyển khoản vị thế | `77E` = `GIVEUP`/`TAKEUP` (MT598 sub 631) |
| ACK / NAK | Điện xác nhận / từ chối ở mức hệ thống | Xem [02](02-cau-truc-dien-fin.md) mục 1.4 |

## 4. Quy ước tài khoản tiền tại NHTT (BICCODE.x)

Các tag `25`, `50K`, `59` của MT103/MT900/MT910 dùng hậu tố sau BICCODE của TVBT. Nghĩa tổng hợp từ các mục trong tài liệu (diễn giải):

| Hậu tố | Loại tài khoản | Nơi xuất hiện |
|---|---|---|
| `BICCODE.C` | TK tiền **ký quỹ môi giới / khách hàng** (đứng tên VSD) | 2.1, 2.2, 3.6 (MT900/910), 6.1, 6.2 |
| `BICCODE.P` | TK tiền **ký quỹ tự doanh** (đứng tên VSD) | 2.1, 2.2, 6.1, 6.2 |
| `BICCODE.R` | TK tiền gửi TVBT **đã đăng ký với NHTT** (nhận tiền rút/lãi) | 2.2 (tag 59), 2.5 |
| `BICCODE.S` | TK tiền **thanh toán** (thanh toán VM, đáo hạn) của TVBT | 4.2, 5.1.2, 6.1, 6.2 |

Tài khoản ở cấp NĐT/CCP, theo `70E ADTX` của MT598 xác nhận mở tài khoản (mục 1.1): `TA//` = TK giao dịch; `MA//` = TK ký quỹ tại VSD; `CA/` = số định danh tại CCP (ví dụ `CA/VSD00000001`).

Giá trị tag `72`/`70` (diễn giải nội dung) gặp trong các điện tiền: `IM//<số TK ký quỹ>` (nộp/rút ký quỹ), `IM2VM//<số TK ký quỹ NĐT>` (rút ký quỹ để thanh toán VM), `AT//<TK chuyển>/<TK nhận>` (chuyển khoản tất toán), `CRED` / `SETL` / `PAYMENT` (thanh toán VM), `DELI` (thanh toán đáo hạn), `TTDM` (thanh toán thay thế bằng tiền khi mất khả năng chuyển giao).

## 5. Quy ước định dạng trường (cột Content trong các bảng điện)

Theo chuẩn SWIFT ISO 15022 (diễn giải):

| Ký hiệu | Nghĩa |
|---|---|
| `n` / `a` / `c` / `x` | số / chữ hoa / chữ-số / ký tự bất kỳ (trong bộ ký tự hợp lệ) |
| `d` | số thập phân |
| `e` | khoảng trắng |
| `3!n` | đúng 3 chữ số |
| `16x` | tối đa 16 ký tự |
| `4*35x` | tối đa 4 dòng, mỗi dòng tối đa 35 ký tự |
| `[ ]` | thành phần tùy chọn |
| `:4!c//8!n` | qualifier 4 ký tự + `//` + 8 chữ số (ví dụ `:PREP//20160323`) |

Cột **Status**: `M` = bắt buộc (Mandatory), `O` = tùy chọn (Optional). Các dòng `16R`/`16S` là mở/đóng khối con (Start/End of Block) trong điện dạng sequence.
