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
		id_remove, _ := cmd.Flags().GetUint("remove")
		if verbose {
			fmt.Println("[ID] ", id_remove)
		}
		core.RemoveTask(id_remove, rows, verbose)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	rootCmd.Flags().UintP("remove", "d", 0, "Removing a task via an ID")
}
