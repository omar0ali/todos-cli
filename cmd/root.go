/*
Copyright © 2025 OMAR BAGUNAID BAJUNAIDOMAR@GMAIL.COM
*/
package cmd

import (
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
	Use:     "todos",
	Short:   "A simple task manager and to-do list tracker.",
	Long:    "Todos is a CLI application that helps users manage their to-do tasks. You can add, remove, edit tasks, assign statuses, and display the latest tasks.",
	Version: "1.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.PersistentFlags().GetBool("verbose")
		if verbose {
			log.Println("[VERBOSE] ON")
			utils.GetCurrentDir(verbose)
		}
		if err := cmd.Help(); err != nil {
			log.Println("Error showing help:", err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	defer utils.SaveData("tasks.json", rows)
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Show debug messages")
	utils.LoadingData("tasks.json", &rows)
}
