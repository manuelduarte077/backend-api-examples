package cmd

import (
	"fmt"
	"strconv"
	"todolist/tasks"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		err = tasks.DeleteTask(id)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}

		fmt.Println("Task deleted:", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
