package cmd

import (
	"github.com/omar0ali/todos/core"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a selected task",
	Long:  "Edit a selected task's title and description",
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		id, _ := cmd.Flags().GetUint("id")
		title, _ := cmd.Flags().GetString("title")
		desc, _ := cmd.Flags().GetString("desc")
		core.EditTask(rows, id, title, desc, verbose)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().UintP("id", "i", 0, "The id of a task to edit")
	editCmd.Flags().StringP("title", "t", "", "The title of a task to edit")
	editCmd.Flags().StringP("desc", "d", "", "The description of a taesk to edit")
}
