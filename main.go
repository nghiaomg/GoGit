package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

const version = "1.0.0"

// executeGitCommand thực thi lệnh git và trả về kết quả
func executeGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// checkGitRepo kiểm tra xem thư mục hiện tại có phải là git repository không
func checkGitRepo() error {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	return cmd.Run()
}

// pushCommand thực hiện git add, commit và push
func pushCommand(cmd *cobra.Command, args []string) error {
	// Kiểm tra xem có phải git repository không
	if err := checkGitRepo(); err != nil {
		return fmt.Errorf("không phải git repository. Hãy chạy 'git init' trước")
	}

	// Kiểm tra commit message
	if len(args) == 0 {
		return fmt.Errorf("vui lòng cung cấp commit message")
	}

	commitMessage := strings.Join(args, " ")

	fmt.Println("🚀 Bắt đầu quá trình git add, commit và push...")

	// Bước 1: git add .
	fmt.Println("📁 Đang thêm tất cả file thay đổi...")
	if err := executeGitCommand("add", "."); err != nil {
		return fmt.Errorf("lỗi khi thực hiện 'git add .': %v", err)
	}
	fmt.Println("✅ Đã thêm tất cả file thay đổi")

	// Bước 2: git commit
	fmt.Printf("💬 Đang commit với message: '%s'...\n", commitMessage)
	if err := executeGitCommand("commit", "-m", commitMessage); err != nil {
		return fmt.Errorf("lỗi khi thực hiện 'git commit': %v", err)
	}
	fmt.Println("✅ Đã commit thành công")

	// Bước 3: git push
	fmt.Println("🌐 Đang push lên remote repository...")
	if err := executeGitCommand("push"); err != nil {
		return fmt.Errorf("lỗi khi thực hiện 'git push': %v", err)
	}
	fmt.Println("🎉 Đã push thành công!")

	return nil
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ggit",
		Short: "GGit - Go Git CLI Tool",
		Long: `GGit là một công cụ CLI đơn giản để rút gọn các lệnh Git thường dùng.
Thay vì chạy 3 lệnh riêng biệt (git add, commit, push), 
bạn chỉ cần chạy 1 lệnh duy nhất.`,
	}

	var pushCmd = &cobra.Command{
		Use:   "push [commit message]",
		Short: "Thực hiện git add, commit và push trong một lệnh",
		Long: `Lệnh push sẽ thực hiện tuần tự:
1. git add . (thêm tất cả file thay đổi)
2. git commit -m "commit message"
3. git push

Ví dụ: ggit push "fix bug login"`,
		RunE: pushCommand,
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Hiển thị phiên bản của ggit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("ggit version %s\n", version)
		},
	}

	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Lỗi: %v\n", err)
		os.Exit(1)
	}
} 