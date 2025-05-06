package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task",
	Long:  "Adding tasks",
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		title, _ := cmd.Flags().GetString("title")
		desc, _ := cmd.Flags().GetString("desc")
		if verbose {
			fmt.Println("Verbose IS ON SO WE ADDING HERE TOO")
			fmt.Println("You have used the title: ", title, desc)
			fmt.Println("add called")
		}
		if len(args) > 0 {
			fmt.Println("OO, we have args now :P")
			fmt.Println(args)
		}
		fmt.Println("The Info: ", rows)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("title", "t", "untilte", "Name the task")
	addCmd.Flags().StringP("desc", "d", "description", "Description of the task")
}
