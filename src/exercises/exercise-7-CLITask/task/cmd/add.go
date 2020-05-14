package cmd

import (
	"exercises/exercise-7-CLITask/task/db"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.TrimSpace(strings.Join(args, " "))
		_, err := db.AddTask(task)
		if err != nil {
			fmt.Println("Error ", err.Error())
		} else {
			fmt.Println("Added \"" + task + "\" to your task list")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
