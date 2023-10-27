package cmd

import (
	"fmt"

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
	rootCmd.SetHelpFunc(timeTravelHelp)
	rootCmd.AddCommand(toCmd)
	rootCmd.AddCommand(feedbackCmd)
}

func timeTravelHelp(cmd *cobra.Command, args []string) {
	helpMessage := `
    timetravel - Time Travel CLI

    Usage:
      timetravel [command] [arguments]

    Commands:
      to        Travel to a specified daytime (in MoDDYYYYHHMi format).
                Example: 
                  timetravel to 070417761200 
                  This will time travel to July 4, 1776 at 12:00am.

                Options:
                  -o, --output    Specify the output format. 
                                  Valid formats: text (default), json, yaml, dashboard.
                                  Example:
                                    timetravel to 070417761200 -o=json
                                    timetravel to 070417761200 --output=json

      feedback  Post travel survey to gather your experience.
                Example:
                  timetravel feedback

                Options:
                  --no-input      Run the survey without user input. 

    Options:
      -h, --help       Show this help message and exit.

    Note:
      The 'dashboard' output format is a terminal dashboard meant for human consumption only.

    Need more details or facing issues? Refer to the official documentation at [official_documentation_link].
    `
	fmt.Println(helpMessage)
}
