package cmd

import (
	"fmt"
	"strconv"
	"todolist/tasks"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task ID]",
	Short: "Mark a task as completed",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		err = tasks.CompleteTask(id)
		if err != nil {
			fmt.Println("Error completing task:", err)
			return
		}

		fmt.Println("Task completed:", id)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
