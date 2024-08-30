package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/lloydrichards/task/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "CLI task manager",
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printTable(tasks []db.Task) {
	w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', tabwriter.Debug)
	fmt.Print("\n")
	fmt.Fprintln(w, "ID\t DESCRIPTION\t STATUS\t CREATED")
	for _, task := range tasks {
		fmt.Fprint(w, task.ID, "\t ", task.Description, "\t ", task.Status, "\t ", task.Created.Format("15:04"), "\n")
	}
	fmt.Print("\n")
	w.Flush()
}
