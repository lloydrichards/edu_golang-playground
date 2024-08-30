package cmd

import (
	"strconv"

	"github.com/lloydrichards/task/db"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task on your task list",
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
		task, err := t.GetTask(id)
		if err != nil {
			return err
		}

		desc, err := cmd.Flags().GetString("description")
		if desc == "" {
			desc = task.Description
		}
		if err != nil {
			return err
		}
		prog, err := cmd.Flags().GetInt("status")
		if err != nil {
			return err
		}

		if err := t.UpdateTask(id, desc, prog); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP(
		"description",
		"d",
		"",
		"specify a description for your task",
	)
	updateCmd.Flags().IntP(
		"status",
		"s",
		0,
		"specify a status for your task",
	)

}
