package cmd

import (
	"fmt"

	"github.com/lloydrichards/task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := db.OpenDB(db.SetupPath())
		if err != nil {
			return err
		}
		defer t.DB.Close()
		tasks, err := t.GetTasks()
		if err != nil {
			return err
		}
		fmt.Print(tasks)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
