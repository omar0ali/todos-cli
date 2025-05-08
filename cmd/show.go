package cmd

import (
	"fmt"
	"github.com/omar0ali/todos/core"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Showing the list of tasks",
	Long:  "Showing a list of tasks in a table, which also includes the status of each task.",
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		status, _ := cmd.Flags().GetString("status")
		if status == "" {
			core.DisplayRecords(rows, verbose)
		} else {
			if core.IsValidStatus(core.Status(status)) {
				core.DisplayRecordsStatus(rows, verbose, core.Status(status))
			} else {
				fmt.Println("Unknown status:", status)
				fmt.Println("Valid options: todo, in-progress, done")
				core.DisplayRecords(rows, verbose)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("status", "s", "todo", "Display a list of tasks based on its status")
}
