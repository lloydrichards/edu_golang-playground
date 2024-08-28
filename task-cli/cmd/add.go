package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your task list",
	Long:  `Add a new task to your task list`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("add called")
		// Do Stuff Here
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
