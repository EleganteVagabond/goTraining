package cmd

import (
	"exercises/exercise-7-CLITask/task/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes task(s), by number from list command, from your list",
	Run: func(cmd *cobra.Command, args []string) {
		ids := make([]int, len(args))
		for ix, arg := range args {
			intval, err := strconv.Atoi(arg)
			if err == nil {
				ids[ix] = intval
			} else {
				fmt.Println(arg, "is not an properly formed number (i.e. 1,2,3) ")
			}
		}
		removed, err := db.RemoveTasks(ids)
		if err == nil {
			// some ids not removed
			if len(removed) != len(ids) {
				badIDs := make([]int, 0)
				rem := make(map[int]bool, len(ids)-len(removed))
				for _, id := range removed {
					rem[id] = true
				}
				for _, id := range ids {
					if !rem[id] {
						badIDs = append(badIDs, id)
					}
				}
				fmt.Println("The following tasks were not found", badIDs)
			}
			if len(removed) != 0 {
				fmt.Println("Task(s)", removed, "removed. Don't worry, we won't tell nobody")
			}
		} else {
			fmt.Println("Error occurred while doing (imagine that :) )", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
