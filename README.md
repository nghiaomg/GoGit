# GGit - Go Git CLI Tool

GGit là một công cụ CLI đơn giản được viết bằng Go để rút gọn các lệnh Git thường dùng.

**Tác giả:** nghiaomg

> **🚨 DISCLAIMER: GGit là công cụ hỗ trợ, KHÔNG THAY THẾ Git!**  
> Bạn vẫn cần cài đặt Git và hiểu cách sử dụng Git cơ bản. GGit chỉ giúp gõ lệnh nhanh hơn.

## Mô tả

GGit giúp rút gọn các lệnh Git phổ biến thành các lệnh đơn giản hơn:

### Khởi tạo repository
Thay vì chạy 6 lệnh:
```bash
git init
git add .
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/user/repo.git
git push -u origin main
```

Chỉ cần 1 lệnh:
```bash
ggit init https://github.com/user/repo.git
```

### Push thay đổi
Thay vì chạy 3 lệnh:
```bash
git add .
git commit -m "commit message"
git push
```

Chỉ cần 1 lệnh:
```bash
ggit push "commit message"
```

## Cài đặt

### Yêu cầu trước khi cài đặt
- Git đã được cài đặt và cấu hình
- Go 1.19 trở lên

### Từ source code
```bash
git clone https://github.com/nghiaomg/GoGit.git
cd GoGit
go build -o ggit.exe main.go
```

### Sử dụng go install (nếu có)
```bash
go install github.com/nghiaomg/GoGit@latest
```

## Sử dụng

### Khởi tạo repository mới
```bash
# Khởi tạo repository với commit message mặc định
ggit init https://github.com/user/repo.git

# Khởi tạo repository với commit message tùy chỉnh
ggit init https://github.com/user/repo.git "initial project setup"
```

### Push thay đổi
```bash
# Thêm tất cả file, commit và push
ggit push "your commit message"
```

### Các lệnh khác
```bash
# Hiển thị trạng thái repository
ggit status

# Pull từ remote repository
ggit pull

# Clone repository
ggit clone https://github.com/nghiaomg/GoGit

# Hiển thị trợ giúp
ggit help

# Hiển thị phiên bản
ggit version
```

## Yêu cầu

- Go 1.19 trở lên
- Git đã được cài đặt và cấu hình
- Repository Git đã được khởi tạo

## ⚠️ Lưu ý quan trọng

- **GGit KHÔNG THAY THẾ Git**: Đây chỉ là công cụ wrapper để viết lệnh nhanh hơn
- **Vẫn cần Git**: Bạn vẫn cần cài đặt và cấu hình Git trên máy
- **Chỉ là shortcut**: GGit chỉ gói gọn các lệnh Git phổ biến, không có tính năng mới
- **Git credentials**: Đảm bảo bạn đã cấu hình Git credentials (username, email, SSH keys) trước khi sử dụng
- **Thêm tất cả file**: Tool sẽ thực hiện `git add .` (thêm tất cả file thay đổi)
- **Kiểm tra trước khi dùng**: Luôn kiểm tra `git status` trước khi push để tránh commit nhầm file

## Giấy phép

MIT License 