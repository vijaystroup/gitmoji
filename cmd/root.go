package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gitm",
	Version: "0.1.0",
	Short:   "😎 Gitmoji formats your git commit messages.",
	Long: `😎 Gitmoji is an extensible git commit message formatter.
It prepends an emoji to your git commit messages.
View this project and documentation at https://www.github.com/VijayStroup/gitmoji.
Gitmoji v0.1.0 licensed under the Apache License, Version 2.0.`,
	Example: `gitm new I just made this super awesome new addition to my project!
gitm fix This thing shouldn't be acting up anymore.
gitm update Added on error checking to this function.`,
	SilenceErrors: true, // silence errors of commands to use custom error text
	SilenceUsage:  true, // silence errors of flags to use custom error text
}

// Execute of root command 'gitm'
// Check for invalid commands/flags and display the error along with the help text
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		errorMessage := err.Error()
		errorMessage = strings.Title(string(errorMessage[0])) + errorMessage[1:]

		// the craziness going on between the %s. is changing the color to red
		// and then changing the color back to the default for Help()
		fmt.Printf("\033[31m%s.\033[0m\n\n", errorMessage)
		rootCmd.Help()
	}
}
