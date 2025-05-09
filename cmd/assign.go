package cmd

import (
	"log"

	"github.com/omar0ali/todos/core"
	"github.com/spf13/cobra"
)

var assignCmd = &cobra.Command{
	Use:   "assign",
	Short: "Assign a new status to a task by its ID",
	Long: `Assign a new status to a specific task using its ID.
Available status options: todo, in-progress, or done.
Example: todos assign -d 1 -s in-progress`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		id, _ := cmd.Flags().GetUint("id")
		status, _ := cmd.Flags().GetString("status")

		if verbose {
			log.Printf("[ASSIGN]")
			log.Printf("[DATA] Task ID: %d", id)
			log.Printf("[DATA] New Status: %s", status)
		}

		if id == 0 {
			log.Println("[Error]: Task ID is required (use -d or --id)")
			return
		}

		if status == "" {
			log.Println("[Error]: Status must be one of: todo, in-progress, done (use -s or --status)")
			return
		}

		if core.IsValidStatus(core.Status(status)) {
			core.UpdateStatusTask(id, core.Status(status), rows, verbose)
		} else {
			log.Println("Unknown status:", status)
			log.Println("Valid options: todo, in-progress, done")
		}
	},
}

func init() {
	rootCmd.AddCommand(assignCmd)
	assignCmd.Flags().UintP("id", "d", 0, "ID of the task to update")
	assignCmd.Flags().StringP("status", "s", "todo", "New status to assign (todo, in-progress, done)")
}
