package cmd

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/fatih/color"
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
	header := color.New(color.FgHiMagenta)
	command := color.New(color.FgHiCyan)
	option := color.New(color.FgHiGreen)
	link := color.New(color.FgBlue)
	example := color.New(color.FgWhite)

	helpMessage := "\n"
	helpMessage += fmt.Sprintf("Timetravel CLI %v%v\n\n", emoji.Automobile, emoji.HighVoltage)
	helpMessage += header.Sprint("usage:\n")
	helpMessage += "  " + command.Sprint("timetravel") + " [command] [arguments]\n\n"
	helpMessage += header.Sprint("commands:\n")
	helpMessage += "  " + command.Sprint("to") + "        Travel to a specified daytime (in MoDDYYYYHHMi format).\n"
	helpMessage += "            Example:\n"
	helpMessage += "              " + example.Sprint("timetravel to 102119850120") + "\n"
	helpMessage += "            This will time travel to BTTF Day - Oct 21, 1985 at 01:20am.\n\n"
	helpMessage += "            Options:\n"
	helpMessage += "              " + option.Sprint("-o, --output") + "    Specify the output format. \n"
	helpMessage += "                              Valid formats: text (default), json, yaml, dashboard.\n"
	helpMessage += "                              Example:\n"
	helpMessage += "                                " + example.Sprint("timetravel to 070417761200 -o=json") + "\n"
	helpMessage += "                                " + example.Sprint("timetravel to 070417761200 --output=json") + "\n\n"
	helpMessage += "            Note:\n"
	helpMessage += "              " + "The 'dashboard' output format is a terminal dashboard meant for human consumption only.\n\n"
	helpMessage += "  " + command.Sprint("feedback") + "  Post travel survey to gather your experience.\n"
	helpMessage += "            Example:\n"
	helpMessage += "              " + example.Sprint("timetravel feedback") + "\n\n"
	helpMessage += "            Options:\n"
	helpMessage += "              " + option.Sprint("--no-input") + "      Run the survey without user input. \n\n"
	helpMessage += header.Sprint("options:\n")
	helpMessage += "  " + option.Sprint("-h, --help") + "       Show this help message and exit.\n\n"
	helpMessage += "Need more details or facing issues? Refer to the official documentation at " + link.Sprint("http://timetravelcli.com") + "\n"
	fmt.Println(helpMessage)
}
