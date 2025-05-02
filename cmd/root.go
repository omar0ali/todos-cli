/*
Copyright © 2025 OMAR BAGUNAID BAJUNAIDOMAR@GMAIL.COM
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "todos",
	Short: "a simple taks manager or todo list tracker.",
	Long: `
todos is a CLI application that help users to manager a todo list tasks. We can add remove edit tasks as well as assign completed tasks and show latest tasks.
`,
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
