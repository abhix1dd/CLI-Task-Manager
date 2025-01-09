package cmd

import (
    "cli-task-manager/storage"
    "fmt"

    "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        filePath := "tasks.json"
        tasks, _ := storage.LoadTasks(filePath)

        for _, task := range tasks {
            status := "Pending"
            if task.Completed {
                status = "Completed"
            }
            fmt.Printf("%d. %s [%s]\n", task.ID, task.Description, status)
        }
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}
