/*
Copyright © 2025 OMAR BAGUNAID BAJUNAIDOMAR@GMAIL.COM
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	data string
)

var rootCmd = &cobra.Command{
	Use:   "todos",
	Short: "a simple taks manager or todo list tracker.",
	Long: `
todos is a CLI application that help users to manager a todo list tasks. We can add remove edit tasks as well as assign completed tasks and show latest tasks.
`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.PersistentFlags().GetBool("verbose")
		if verbose {
			fmt.Println("Verbose is ON")
		}
		// Loading Phase
		fileInfo, err := os.Stat("dat.json")
		if err != nil {
			f, err := os.Create("dat.json")
			if err != nil {
				log.Panic(err)
			}
			defer f.Close()
		}

		fmt.Println("We are passing")
		fmt.Println(fileInfo)

		data, err := os.ReadFile("dat.json")
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(data)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "show debug messages")
}
