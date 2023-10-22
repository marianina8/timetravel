package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "timeTravel",
	Short: "A CLI tool for time travel!",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(travelCmd)
	rootCmd.AddCommand(surveyCmd)
}
