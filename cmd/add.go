package cmd

import (
    "cli-task-manager/models"
    "cli-task-manager/storage"
    "fmt"

    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new task",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        taskDescription := args[0]
        filePath := "tasks.json"
        tasks, _ := storage.LoadTasks(filePath)

        newTask := models.Task{
            ID:          len(tasks) + 1,
            Description: taskDescription,
            Completed:   false,
        }
        tasks = append(tasks, newTask)
        storage.SaveTasks(filePath, tasks)

        fmt.Printf("Task added: %s\n", taskDescription)
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}
