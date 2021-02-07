package tests

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

// setup test directory with git to test commit messages
func init() {
	// make mock directory to initialize git repo
	os.Mkdir("mock_project", 0777)
	if err := os.Chdir("mock_project"); err != nil {
		panic(fmt.Errorf("Unable to change directory: %s", err.Error()))
	}

	// initialize git repo
	cmd := exec.Command("git", "init")
	if err := cmd.Run(); err != nil {
		panic(fmt.Errorf("Unable to initialize git repo: %s", err.Error()))
	}

	// create new file
	editFile()

	// create first commit
	cmd = exec.Command("git", "add", "-A")
	if err := cmd.Run(); err != nil {
		panic(fmt.Errorf("Unable to stage mock file: %s", err.Error()))
	}

	cmd = exec.Command("git", "commit", "-a", "-m", "\"first commit\"")
	if err := cmd.Run(); err != nil {
		panic(fmt.Errorf("Unable to make first commit: %s", err.Error()))
	}
}

func TestCommitDefault(t *testing.T) {
	fmt.Println("IN TEST")
	return
}

// editFile appends to a test file to simulate changes to a file to allow git
// to detect changes for commits
func editFile() {
	f, err := os.OpenFile("mock", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(fmt.Errorf("Unable to create mock file: %s", err.Error()))
	}
	defer f.Close()

	if _, err := f.WriteString(time.Now().String() + "\n"); err != nil {
		panic(fmt.Errorf("Unable to append to mock file: %s", err.Error()))
	}
}
