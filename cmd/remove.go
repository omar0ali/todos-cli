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
		id_remove, _ := cmd.Flags().GetUint("id")
		fmt.Println(id_remove)
		if verbose {
			fmt.Printf("[ID] %d [REMOVED]\n", id_remove)
		}
		core.RemoveTask(id_remove, &rows, verbose)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().UintP("id", "d", 0, "Remove a task by its ID")
}
