package cmd

import (
	"exercises/exercise-7-CLITask/task/db"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Lists the tasks you have completed today",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetCompletedTasks()
		if err != nil {
			panic(err)
		}
		if len(tasks) == 0 {
			fmt.Println("Hey, wait a minute. You haven't completed any tasks!")
		} else {
			now := time.Now()
			todayTasks := make([]db.TaskData, 0)
			for _, task := range tasks {
				if task.Value.UpdateTime.Day() == now.Day() {
					todayTasks = append(todayTasks, task.Value)
				}
			}
			todayTasksCnt := len(todayTasks)
			if todayTasksCnt != 0 {
				fmt.Println("You have finished the following tasks today:")
				for _, task := range todayTasks {
					fmt.Println("-", task.Value)
				}
				if todayTasksCnt < 5 {
					fmt.Println("Not bad but I bet you can do better!")
				} else {
					fmt.Println("How did you do so many!?")
				}
			} else {
				fmt.Println("You ain't done shit today, fool... (shakes head)")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)
}
