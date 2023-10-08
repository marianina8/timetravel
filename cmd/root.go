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
	travelCmd.Flags().StringP("destination", "d", "102619850121", "Destination for time travel in MMDDYYYYHHMM format")
	rootCmd.AddCommand(travelCmd)
	rootCmd.AddCommand(surveyCmd)
}
