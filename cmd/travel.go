package cmd

import (
	"fmt"

	"github.com/marianina8/timetravel/dashboard"
	"github.com/marianina8/timetravel/speedometer"
	"github.com/marianina8/timetravel/utils"
	"github.com/spf13/cobra"
)

var travelCmd = &cobra.Command{
	Use:   "travel",
	Short: "Travel to a specified year!",
	Run:   travelTime,
}

func travelTime(cmd *cobra.Command, args []string) {
	destination, err := cmd.Flags().GetString("destination")
	if err != nil {
		fmt.Println("Unable to get destination: ", err)
	}
	parsedDestination, err := utils.ParseDestination(destination)
	if err != nil {
		fmt.Println("Unable to parse destination: ", err)
	}
	formattedDestination := utils.FormatDestination(parsedDestination)

	dashboard.New(formattedDestination)
	speedometer.New()

	fmt.Println("****************************************************************************************************")
	fmt.Println("Time warp completed!  Welcome to ", formattedDestination, "!")
	fmt.Println("****************************************************************************************************")
}
