package cmd

import (
	"fmt"
	"todolist/tasks"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		taskList, err := tasks.ListTasks()
		if err != nil {
			fmt.Println("Error listing tasks:", err)
			return
		}

		for _, task := range taskList {
			status := " "
			if task.Complete {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
