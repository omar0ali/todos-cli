package cmd

import (
	"github.com/omar0ali/todos/core"
	"github.com/spf13/cobra"
	"log"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display the list of tasks",
	Long:  "Displays all tasks in a formatted table, including the status of each task.",
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		status, _ := cmd.Flags().GetString("status")
		if status == "" {
			core.DisplayRecords(rows, verbose)
		} else {
			if core.IsValidStatus(core.Status(status)) {
				core.DisplayRecordsStatus(rows, verbose, core.Status(status))
			} else {
				log.Println("Unknown status:", status)
				log.Println("Valid options: todo, in-progress, done")
				core.DisplayRecords(rows, verbose)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("status", "s", "", "Display a list of tasks based on its status")
}
