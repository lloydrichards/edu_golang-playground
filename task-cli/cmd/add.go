package cmd

import (
	"github.com/lloydrichards/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your task list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := db.OpenDB(db.SetupPath())
		if err != nil {
			return err
		}
		defer t.DB.Close()

		if err := t.CreateTask(args[0]); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
