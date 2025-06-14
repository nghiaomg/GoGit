# GGit - Go Git CLI Tool

GGit là một công cụ CLI đơn giản được viết bằng Go để rút gọn các lệnh Git thường dùng.

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
ggit clone https://github.com/user/repo.git

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