package cmd

import (
	"strconv"

	"github.com/lloydrichards/task/db"
	"github.com/spf13/cobra"
)

var delete = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from your task list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := db.OpenDB(db.SetupPath())
		if err != nil {
			return err
		}
		defer t.DB.Close()

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		if err := t.DeleteTask(id); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(delete)
}
