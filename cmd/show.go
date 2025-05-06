package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Showing the list of tasks",
	Long:  "Showing a list of tasks in a table, which also includes the status of each task.",
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			for i := 0; i < len(rows); i++ {
				fmt.Println(rows[i])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
