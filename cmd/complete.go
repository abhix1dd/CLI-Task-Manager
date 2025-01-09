package cmd

import (
    "cli-task-manager/storage"
    "fmt"
    "strconv"

    "github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
    Use:   "complete",
    Short: "Mark a task as completed by ID",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        taskID, _ := strconv.Atoi(args[0])
        filePath := "tasks.json"
        tasks, _ := storage.LoadTasks(filePath)

        for i, task := range tasks {
            if task.ID == taskID {
                tasks[i].Completed = true
            }
        }

        storage.SaveTasks(filePath, tasks)
        fmt.Printf("Task with ID %d marked as completed\n", taskID)
    },
}

func init() {
    rootCmd.AddCommand(completeCmd)
}
