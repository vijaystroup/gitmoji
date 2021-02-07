package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// flags
var (
	autoStage bool
)

func init() {
	commands := make(map[string]string)

	// initialize default commands
	commands["new"] = "✨"
	commands["fix"] = "🔧"
	commands["update"] = "☝️"

	// get commands from environment variables
	getEnvs(commands)

	// loop through commands and add them to rootCmd
	for k, v := range commands {
		rootCmd.AddCommand(makeCommand(k, v))
	}
}

// makeCommand generates a pointer to a cobra.Command to add to the rootCmd.
// This is done dynamically in the init() function.
func makeCommand(name, emoji string) *cobra.Command {
	command := &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("Prepend %s to git commit message", emoji),
		Run: func(cmd *cobra.Command, args []string) {
			message := emoji + " " + strings.Join(args, " ")

			// differnet commits dependent on flag
			var commit *exec.Cmd
			if autoStage == true {
				// track untracked files
				if err := exec.Command("git", "add", "-A").Run(); err != nil {
					fmt.Println("\033[31mgitm: 🚨 No git repo found in the current directory.\033[0m")
					os.Exit(1)
				}
				commit = exec.Command("git", "commit", "-am", message)
			} else {
				commit = exec.Command("git", "commit", "-m", message)
			}

			// run git command
			if err := commit.Run(); err != nil {
				// if the exitcode is 128, that is an indication of the current
				// directory not being a git repo, so let's tell the user
				if exitError, _ := err.(*exec.ExitError); exitError.ExitCode() == 128 {
					fmt.Println("\033[31mgitm: 🚨 No git repo found in the current directory.\033[0m")
					os.Exit(2)
				}

				// check for untracked files that could be commited
				untracked := exec.Command("git", "ls-files", "--other", "--exclude-standard", "--directory")
				untrackedOut, _ := untracked.Output()
				if len(untrackedOut) != 0 {
					fmt.Println("\033[31mgitm: Files available for staging but none found for commit.")
					fmt.Println("Try 'git add FILE' or 'gitm COMMAND -a MESSAGE'.\033[0m")
					os.Exit(3)
				}

				// check git diff to see if repo is up to date
				diff := exec.Command("git", "diff")
				diffOut, _ := diff.Output()
				if len(diffOut) == 0 {
					fmt.Println("\033[33mgitm: Nothing to commit, working tree clean.\033[0m")
					os.Exit(0)
				} else {
					fmt.Println("\033[31mgitm: Modified files not staged for commit.\033[0m")
					os.Exit(4)
				}

				// unknown error
				fmt.Println("\033[31mgitm: Unable to run command.\033[0m")
				os.Exit(5)
			}

			// if no error, report the commit message
			fmt.Printf("\033[32mSuccessfully commited:\033[0m %s\n", message)
		},
	}

	// flags
	command.Flags().BoolVarP(&autoStage, "all", "a", false,
		"Automatically track and stage all files to commit")

	return command
}

// getEnvs adds to a map of environment variables set by the user specific to
// Gitmoji. They are then processed to ensure they meet the correct format of
// "command:emoji". Example: "fix:🔧"
func getEnvs(c map[string]string) {
	for _, env := range os.Environ() {
		// check if env is not a Gitmoji env
		if !strings.HasPrefix(env, "GITM_") {
			continue
		}

		// get env name and value
		envSplit := strings.Split(env, "=")
		name := envSplit[0]
		value := strings.ReplaceAll(envSplit[1], " ", "")

		// check to see if value format is good
		v := strings.Split(value, ":")
		if len(v) != 2 {
			fmt.Printf("\033[31mEnvironment variable '%s' is is of wrong format: '%s'.\n", name, value)
			fmt.Printf("The correct format is 'command:emoji'.\nExample: 'fix:🔧'.\033[0m\n\n")
			os.Exit(6)
		}

		// check to see if name is lowercase
		if v[0] != strings.ToLower(v[0]) {
			fmt.Printf("\033[33mWarning: environment variable %s does not have key of type lowercase.\n", name)
			fmt.Printf("'%s' will be treated as '%s'. Change %s to supress this warning.\033[0m\n\n", v[0], strings.ToLower(v[0]), name)
		}
		v[0] = strings.ToLower(v[0])

		// all checks passed, add to commands map
		c[v[0]] = v[1]
	}
}
