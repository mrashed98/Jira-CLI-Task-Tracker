/*
Copyright © 2024 MahmoudRashed mahmoudrashed2806@gmail.com
*/
package cmd

import (
	"fmt"

	jira "github.com/mrashed98/jiraCliTracker/jira"
	"github.com/spf13/cobra"
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Get all tasks",
	Long:  `Get all Tasks assigned to you with the specific label your configured in the env file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tasks called")
		jira.GetAllTasks()
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)
}
