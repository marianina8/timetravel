package cmd

import (
	"github.com/marianina8/timetravel/survey"
	"github.com/spf13/cobra"
)

var surveyCmd = &cobra.Command{
	Use:   "survey",
	Short: "Give feedback on your time travel experience!",
	Run: func(cmd *cobra.Command, args []string) {
		survey.Run()
	},
}
