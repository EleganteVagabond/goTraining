package cmd

import (
	"exercises/exercise-7-CLITask/task/db"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the tasks in your task list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetIncompleteTasks()
		if err != nil {
			panic(err)
		}
		if len(tasks) == 0 {
			fmt.Println("Hooray, no tasks found! Why not have a nap?")
		}
		for ix, task := range tasks {
			fmt.Println(ix+1, "-", task.Value.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
