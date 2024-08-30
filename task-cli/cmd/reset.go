package cmd

import (
	"github.com/lloydrichards/task/db"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Delete all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := db.DeleteDB(db.SetupPath())
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
