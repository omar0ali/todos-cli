package cmd

import (
	"fmt"
	"github.com/omar0ali/todos/core"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "removing a task from a list",
	Long:  "removing a task from tasks list by providing an id of the task",
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
	removeCmd.Flags().UintP("id", "d", 0, "Removing a task via an ID")
}
