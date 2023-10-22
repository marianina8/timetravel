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
	if len(args) == 0 {
		fmt.Println("Please enter a destination time!  The format input is MonthDayYearHourMinute -> MoDDYYYYHHMi")
		return
	}
	destination := args[0]
	parsedDestination, err := utils.ParseDestination(destination)
	if err != nil {
		fmt.Println("Unable to parse destination: ", err)
		fmt.Println("For example, to travel to 12/25/2020 at 12:00am, enter: MoDDYYYYHHMi -> 122520200000")
	}
	formattedDestination := utils.FormatDestination(parsedDestination)
	if dashboard.New(formattedDestination) {
		speedometer.New()
		utils.Print("Time warp completed!  Welcome to " + formattedDestination + "!")
		return
	}
	utils.Print("Flux Capacitor Deactivated: Stay in the present, McFly!")
}
