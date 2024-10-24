package cmd

import (
	"fmt"
	"todolist/tasks"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		err := tasks.AddTask(task)
		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}
		fmt.Println("Task added:", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
