/*
Copyright © 2024 Mahmoud Rashed mahmoudrashed2806@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jiraCliTracker",
	Short: "Jira CLI Task Tracker Application",
	Long: `This apps Aims to bring Tracking your tasks on Jira to the CLI
	
	You can check all the tasks assigned to you with couple of filters (not all implemented yet)
	- Tasks with specific label (current default)
	- Tasks with specific project (not yet implemented)
	- Tasks with specific sprint (not yet implemented)
	- Tasks with specific Date range (not yet implemented)
	- Tasks with sepcific Status (not yet implemented)
	- A lot more comming ...

	And you can also update the status for the task (not yet implemented)
	`,
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
