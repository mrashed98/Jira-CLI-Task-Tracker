/*
Copyright © 2024 MahmoudRashed mahmoudrashed2806@gmail.com
*/
package cmd

import (
	jira "github.com/mrashed98/jiraCliTracker/jira"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "This Command Returns all the tasks marked as Done",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		jira.GetCompletedTasks()
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
