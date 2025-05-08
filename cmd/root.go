/*
Copyright © 2025 OMAR BAGUNAID BAJUNAIDOMAR@GMAIL.COM
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/omar0ali/todos/core"
	"github.com/omar0ali/todos/utils"
	"github.com/spf13/cobra"
)

var (
	rows []core.TableCl
)

var rootCmd = &cobra.Command{
	Use:   "todos",
	Short: "a simple taks manager or todo list tracker.",
	Long:  "todos is a CLI application that help users to manager a todo list tasks. We can add remove edit tasks as well as assign completed tasks and show latest tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.PersistentFlags().GetBool("verbose")
		if verbose {
			log.Println("[VERBOSE] ON")
			utils.GetCurrentDir(verbose)
		}
		fmt.Println(cmd.Help())
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	defer utils.SaveData("tasks.json", rows) //ensure this always execute at the end
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "show debug messages")
	utils.LoadingData("tasks.json", &rows)
}
