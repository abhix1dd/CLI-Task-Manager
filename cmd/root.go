package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-task-manager",
	Short: "A CLI task manager application",
	Long: `cli-task-manager is a command-line tool to help manage tasks efficiently.
It supports adding, deleting, marking tasks as completed, and listing them.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the CLI Task Manager! Use --help to see available commands.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// It is called by the main function.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
