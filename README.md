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
Thay vì chạy 3-5 lệnh:
```bash
git add .
git commit -m "commit message"
git checkout -b feature-branch  # (nếu cần)
git push -u origin feature-branch
```

Chỉ cần 1 lệnh:
```bash
ggit push "commit message" --branch feature-branch
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
# Push lên branch hiện tại
ggit push "your commit message"

# Push lên branch cụ thể (tạo mới nếu chưa có)
ggit push "update feature" --branch dev
ggit push "hotfix" -b hotfix/critical-bug
```

### Quản lý branch
```bash
# Push lên branch hiện tại
ggit push "fix bug"

# Tạo branch mới và push
ggit push "new feature" --branch feature/user-auth

# Push lên development branch
ggit push "update docs" -b dev

# Push hotfix
ggit push "critical fix" --branch hotfix/security
```

### Các lệnh khác
```bash
# Hiển thị trạng thái repository
ggit status

# Pull từ remote repository
ggit pull

# Clone repository (full command)
ggit clone https://github.com/nghiaomg/GoGit

# Clone repository (shorthand)
ggit c https://github.com/nghiaomg/GoGit

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