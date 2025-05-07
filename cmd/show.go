package cmd

import (
	"github.com/omar0ali/todos/core"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Showing the list of tasks",
	Long:  "Showing a list of tasks in a table, which also includes the status of each task.",
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		core.DisplayRecords(rows, verbose)

	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
