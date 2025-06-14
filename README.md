# GGit - Go Git CLI Tool

GGit là một công cụ CLI đơn giản được viết bằng Go để rút gọn các lệnh Git thường dùng.

## Mô tả

Thay vì phải chạy 3 lệnh Git riêng biệt:
```bash
git add .
git commit -m "commit message"
git push
```

Bạn chỉ cần chạy 1 lệnh duy nhất:
```bash
ggit push "commit message"
```

## Cài đặt

### Từ source code
```bash
git clone <repository-url>
cd G-Git
go build -o ggit main.go
```

### Sử dụng go install
```bash
go install
```

## Sử dụng

### Lệnh cơ bản
```bash
# Thêm tất cả file, commit và push
ggit push "your commit message"
```

### Các tùy chọn
```bash
# Hiển thị trợ giúp
ggit help

# Hiển thị phiên bản
ggit version
```

## Yêu cầu

- Go 1.19 trở lên
- Git đã được cài đặt và cấu hình
- Repository Git đã được khởi tạo

## Lưu ý

- GGit chỉ là công cụ hỗ trợ viết nhanh hơn, không thay thế Git
- Đảm bảo bạn đã cấu hình Git credentials trước khi sử dụng
- Tool sẽ thực hiện `git add .` (thêm tất cả file thay đổi)

## Giấy phép

MIT License 