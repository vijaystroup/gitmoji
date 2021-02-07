package tests

import (
	"bytes"
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

// TestCommitDefault tests the default commands that come with Gitmoji and
// verifies that the git log reflects the correct messages
func TestCommitDefault(t *testing.T) {
	// setup default commands
	commands := map[string]string{
		"new":    "✨",
		"fix":    "🔧",
		"update": "☝️",
	}

	// go through all commands and test
	for k, v := range commands {
		// edit file and then commit with new command
		editFile()

		msg := "this is a " + k + " commit"
		cmd := exec.Command("go", "run", "../../gitmoji.go", k, "-a", msg)
		if err := cmd.Run(); err != nil {
			panic(fmt.Errorf("Unable to run gitmoji: %s", err.Error()))
		}

		// check last commit
		want := []byte(v + " this is a " + k + " commit")

		out, err := exec.Command("git", "log", "-1", "--pretty=%B").Output()
		if err != nil {
			panic(fmt.Errorf("Unable to run gitmoji: %s", err.Error()))
		} else if bytes.Compare(want, out[:len(out)-2]) != 0 {
			t.Fatalf("Commit failed:\nwant:\t%s\ngot:\t%s", string(want), string(out))
		}
	}

	cleanUp()
}

func TestGetEnvs(t *testing.T) {
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

// cleanUp cleans up the temporary mock_project directory for the tests
func cleanUp() {
	if err := exec.Command("rm", "-rf", "../mock_project").Run(); err != nil {
		panic(fmt.Errorf("Unable to delete mock_project dir: %s", err.Error()))
	}
}
