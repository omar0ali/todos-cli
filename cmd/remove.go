package cmd

import (
	"fmt"

	"github.com/omar0ali/todos/core"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a task from the list",
	Long:  "Remove a task from the task list by providing the task's ID.",
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		idRemove, _ := cmd.Flags().GetUint("id")
		fmt.Println(idRemove)
		if verbose {
			fmt.Printf("[ID] %d [REMOVED]\n", idRemove)
		}
		core.RemoveTask(idRemove, &rows, verbose)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().UintP("id", "i", 0, "Remove a task by its ID")
}
