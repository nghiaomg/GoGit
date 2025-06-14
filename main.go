package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

const version = "1.0.1"

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

// initCommand khởi tạo git repository với remote origin
func initCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("vui lòng cung cấp URL của remote repository")
	}

	remoteURL := args[0]
	commitMessage := "first commit"
	if len(args) > 1 {
		commitMessage = strings.Join(args[1:], " ")
	}

	fmt.Println("🎯 Bắt đầu khởi tạo Git repository...")

	// Bước 1: git init
	fmt.Println("📦 Đang khởi tạo Git repository...")
	if err := executeGitCommand("init"); err != nil {
		return fmt.Errorf("lỗi khi thực hiện 'git init': %v", err)
	}
	fmt.Println("✅ Đã khởi tạo Git repository")

	// Bước 2: git add .
	fmt.Println("📁 Đang thêm tất cả file...")
	if err := executeGitCommand("add", "."); err != nil {
		return fmt.Errorf("lỗi khi thực hiện 'git add .': %v", err)
	}
	fmt.Println("✅ Đã thêm tất cả file")

	// Bước 3: git commit
	fmt.Printf("💬 Đang commit với message: '%s'...\n", commitMessage)
	if err := executeGitCommand("commit", "-m", commitMessage); err != nil {
		return fmt.Errorf("lỗi khi thực hiện 'git commit': %v", err)
	}
	fmt.Println("✅ Đã commit thành công")

	// Bước 4: git branch -M main
	fmt.Println("🌿 Đang đổi tên branch thành 'main'...")
	if err := executeGitCommand("branch", "-M", "main"); err != nil {
		return fmt.Errorf("lỗi khi thực hiện 'git branch -M main': %v", err)
	}
	fmt.Println("✅ Đã đổi tên branch thành 'main'")

	// Bước 5: git remote add origin (hoặc set-url nếu đã tồn tại)
	fmt.Printf("🔗 Đang cấu hình remote origin: %s...\n", remoteURL)

	// Kiểm tra xem remote origin đã tồn tại chưa
	checkCmd := exec.Command("git", "remote", "get-url", "origin")
	if err := checkCmd.Run(); err != nil {
		// Remote origin chưa tồn tại, thêm mới
		if err := executeGitCommand("remote", "add", "origin", remoteURL); err != nil {
			return fmt.Errorf("lỗi khi thực hiện 'git remote add origin': %v", err)
		}
		fmt.Println("✅ Đã thêm remote origin")
	} else {
		// Remote origin đã tồn tại, cập nhật URL
		if err := executeGitCommand("remote", "set-url", "origin", remoteURL); err != nil {
			return fmt.Errorf("lỗi khi thực hiện 'git remote set-url origin': %v", err)
		}
		fmt.Println("✅ Đã cập nhật remote origin")
	}

	// Bước 6: git push -u origin main
	fmt.Println("🚀 Đang push lên remote repository...")
	if err := executeGitCommand("push", "-u", "origin", "main"); err != nil {
		return fmt.Errorf("lỗi khi thực hiện 'git push -u origin main': %v", err)
	}
	fmt.Println("🎉 Đã khởi tạo và push repository thành công!")

	return nil
}

// pushCommand thực hiện git add, commit và push
func pushCommand(cmd *cobra.Command, args []string) error {
	// Kiểm tra xem có phải git repository không
	if err := checkGitRepo(); err != nil {
		return fmt.Errorf("không phải git repository. Hãy chạy 'ggit init <remote-url>' trước")
	}

	// Kiểm tra commit message
	if len(args) == 0 {
		return fmt.Errorf("vui lòng cung cấp commit message")
	}

	commitMessage := strings.Join(args, " ")

	// Lấy branch từ flag
	targetBranch, _ := cmd.Flags().GetString("branch")

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

	// Bước 3: Chuyển branch nếu cần
	if targetBranch != "" {
		fmt.Printf("🌿 Đang chuyển/tạo branch '%s'...\n", targetBranch)

		// Kiểm tra xem branch đã tồn tại chưa
		checkBranchCmd := exec.Command("git", "show-ref", "--verify", "--quiet", "refs/heads/"+targetBranch)
		if err := checkBranchCmd.Run(); err != nil {
			// Branch chưa tồn tại, tạo mới và chuyển
			if err := executeGitCommand("checkout", "-b", targetBranch); err != nil {
				return fmt.Errorf("lỗi khi tạo branch '%s': %v", targetBranch, err)
			}
			fmt.Printf("✅ Đã tạo và chuyển sang branch '%s'\n", targetBranch)
		} else {
			// Branch đã tồn tại, chỉ cần chuyển
			if err := executeGitCommand("checkout", targetBranch); err != nil {
				return fmt.Errorf("lỗi khi chuyển sang branch '%s': %v", targetBranch, err)
			}
			fmt.Printf("✅ Đã chuyển sang branch '%s'\n", targetBranch)
		}
	}

	// Bước 4: git push
	if targetBranch != "" {
		fmt.Printf("🌐 Đang push lên branch '%s'...\n", targetBranch)
		if err := executeGitCommand("push", "-u", "origin", targetBranch); err != nil {
			return fmt.Errorf("lỗi khi thực hiện 'git push origin %s': %v", targetBranch, err)
		}
		fmt.Printf("🎉 Đã push thành công lên branch '%s'!\n", targetBranch)
	} else {
		fmt.Println("🌐 Đang push lên remote repository...")
		if err := executeGitCommand("push"); err != nil {
			return fmt.Errorf("lỗi khi thực hiện 'git push': %v", err)
		}
		fmt.Println("🎉 Đã push thành công!")
	}

	return nil
}

// statusCommand hiển thị git status với format đẹp
func statusCommand(cmd *cobra.Command, args []string) error {
	if err := checkGitRepo(); err != nil {
		return fmt.Errorf("không phải git repository")
	}

	fmt.Println("📊 Trạng thái Git repository:")
	return executeGitCommand("status", "--short", "--branch")
}

// pullCommand thực hiện git pull
func pullCommand(cmd *cobra.Command, args []string) error {
	if err := checkGitRepo(); err != nil {
		return fmt.Errorf("không phải git repository")
	}

	fmt.Println("⬇️ Đang pull từ remote repository...")
	if err := executeGitCommand("pull"); err != nil {
		return fmt.Errorf("lỗi khi thực hiện 'git pull': %v", err)
	}
	fmt.Println("✅ Đã pull thành công!")
	return nil
}

// cloneCommand thực hiện git clone và cd vào thư mục
func cloneCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("vui lòng cung cấp URL repository để clone")
	}

	repoURL := args[0]
	fmt.Printf("📥 Đang clone repository: %s...\n", repoURL)

	if err := executeGitCommand("clone", repoURL); err != nil {
		return fmt.Errorf("lỗi khi thực hiện 'git clone': %v", err)
	}

	fmt.Println("✅ Đã clone thành công!")
	fmt.Println("💡 Hãy cd vào thư mục vừa clone để sử dụng các lệnh ggit khác")
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

	var initCmd = &cobra.Command{
		Use:   "init [remote-url] [commit-message]",
		Short: "Khởi tạo Git repository và kết nối với remote",
		Long: `Lệnh init sẽ thực hiện tuần tự:
1. git init (khởi tạo repository)
2. git add . (thêm tất cả file)
3. git commit -m "commit message" (commit với message, mặc định: "first commit")
4. git branch -M main (đổi tên branch thành main)
5. git remote add origin <remote-url> (thêm remote origin)
6. git push -u origin main (push lên remote)

Ví dụ: 
  ggit init https://github.com/user/repo.git
  ggit init https://github.com/user/repo.git "initial project setup"`,
		RunE: initCommand,
	}

	var pushCmd = &cobra.Command{
		Use:   "push [commit message]",
		Short: "Thực hiện git add, commit và push trong một lệnh",
		Long: `Lệnh push sẽ thực hiện tuần tự:
1. git add . (thêm tất cả file thay đổi)
2. git commit -m "commit message"
3. git checkout [branch] (nếu có flag --branch)
4. git push [origin branch]

Ví dụ: 
  ggit push "fix bug login"
  ggit push "update feature" --branch dev
  ggit push "hotfix" -b hotfix`,
		RunE: pushCommand,
	}

	// Thêm flag --branch cho pushCmd
	pushCmd.Flags().StringP("branch", "b", "", "Branch để push (sẽ tạo mới nếu chưa tồn tại)")

	var statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Hiển thị trạng thái Git repository",
		Long:  `Hiển thị trạng thái của Git repository với format ngắn gọn và dễ đọc.`,
		RunE:  statusCommand,
	}

	var pullCmd = &cobra.Command{
		Use:   "pull",
		Short: "Pull từ remote repository",
		Long:  `Thực hiện git pull để cập nhật code từ remote repository.`,
		RunE:  pullCommand,
	}

	var cloneCmd = &cobra.Command{
		Use:   "clone [repository-url]",
		Short: "Clone một repository từ remote",
		Long: `Clone một repository từ remote URL.

Ví dụ: ggit clone https://github.com/user/repo.git`,
		RunE: cloneCommand,
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Hiển thị phiên bản của ggit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("ggit version %s\n", version)
		},
	}

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(cloneCmd)
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Lỗi: %v\n", err)
		os.Exit(1)
	}
}
