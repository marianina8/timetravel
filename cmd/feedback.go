package cmd

import (
	"fmt"

	"github.com/marianina8/timetravel/survey"
	"github.com/spf13/cobra"
)

var feedbackCmd = &cobra.Command{
	Use:   "feedback",
	Short: "Give feedback on your time travel experience!",
	Run: func(cmd *cobra.Command, args []string) {
		skip, err := cmd.Flags().GetBool("no-input")
		if err != nil {
			fmt.Println("Error getting skip flag:", err)
			return
		}
		if skip {
			survey.Skip()
			return
		}
		survey.Run()
	},
}
