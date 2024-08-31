/*
Copyright © 2024 MahmoudRashed mahmoudrashed2806@gmail.com
*/
package cmd

import (
	jira "github.com/mrashed98/jiraCliTracker/jira"
	"github.com/spf13/cobra"
)

// inprogressCmd represents the inprogress command
var inprogressCmd = &cobra.Command{
	Use:   "inprogress",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		jira.GetInProgressTasks()
	},
}

func init() {
	rootCmd.AddCommand(inprogressCmd)
}
