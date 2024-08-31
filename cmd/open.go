/*
Copyright © 2024 MahmoudRashed mahmoudrashed2806@gmail.com
*/
package cmd

import (
	jira "github.com/mrashed98/jiraCliTracker/jira"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "This Command Return all tasks with status Open",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		jira.GetOpenTasks()
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
