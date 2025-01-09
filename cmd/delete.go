package cmd

import (
    "cli-task-manager/models"
    "cli-task-manager/storage"
    "fmt"
    "strconv"

    "github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
    Use:   "delete",
    Short: "Delete a task by ID",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        taskID, _ := strconv.Atoi(args[0])
        filePath := "tasks.json"
        tasks, _ := storage.LoadTasks(filePath)

        updatedTasks := []models.Task{}
        for _, task := range tasks {
            if task.ID != taskID {
                updatedTasks = append(updatedTasks, task)
            }
        }

        storage.SaveTasks(filePath, updatedTasks)
        fmt.Printf("Task with ID %d deleted\n", taskID)
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
}
