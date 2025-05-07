package cmd

import (
	"fmt"

	"github.com/omar0ali/todos/core"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task",
	Long:  "adding tasks",
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		title, _ := cmd.Flags().GetString("title")
		desc, _ := cmd.Flags().GetString("desc")
		status, _ := cmd.Flags().GetString("status")
		if verbose {
			fmt.Println("[ADD]")
			fmt.Println("[VERBOSE ON]")
			fmt.Println("[DATA] ", title)
			fmt.Println("[DATA] ", desc)
			fmt.Println("[DATA] ", status)
		}
		var row core.TableCl
		if core.IsValidStatus(core.Status(status)) {
			row = *core.NewRow(title, desc, core.Status(status))
		} else {
			fmt.Println("Unknown status:", status)
			fmt.Println("Valid options: todo, in-progress, done")
			row = *core.NewRow(title, desc, core.StatusTodo)
		}
		rows = append(rows, row)
		fmt.Println("Data has been added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("title", "t", "untilte", "Name the task")
	addCmd.Flags().StringP("desc", "d", "description", "Description of the task")
	addCmd.Flags().StringP("status", "s", "todo", "Current status of the task")
}
