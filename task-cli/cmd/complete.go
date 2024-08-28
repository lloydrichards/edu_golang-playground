package cmd

import "github.com/spf13/cobra"

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task on your task list as complete",
	Long:  `Mark a task on your task list as complete. You can mark a task as complete by providing the task number.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
