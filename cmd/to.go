package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/marianina8/timetravel/dashboard"
	"github.com/marianina8/timetravel/speedometer"
	"github.com/marianina8/timetravel/utils"
	"github.com/spf13/cobra"

	goyaml "gopkg.in/yaml.v3"
)

var toCmd = &cobra.Command{
	Use:   "to",
	Short: "Go to a specified destination time!",
	Run:   to,
}

func to(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter a destination time!  The format input is MonthDayYearHourMinute -> MoDDYYYYHHMi")
		return
	}
	if len(args) != 1 {
		fmt.Println("Only one parameter expected.  Please enter a destination time in the format input of MonthDayYearHourMinute -> MoDDYYYYHHMi")
		return
	}
	destination := args[0]
	parsedDestination, err := utils.ParseDestination(destination)
	if err != nil {
		fmt.Println("Unable to parse destination: ", err)
		fmt.Println("For example, to travel to 12/25/2020 at 12:00am, enter: MoDDYYYYHHMi -> 122520200000")
	}
	formattedDestination := utils.FormatDestination(parsedDestination)
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		fmt.Println("Error getting output flag:", err)
		return
	}
	switch output {
	case "json":
		err = jsonOutput(formattedDestination)
	case "yaml":
		err = yamlOutput(formattedDestination)
	case "text":
		err = textOutput(formattedDestination)
	case "dashboard":
		if isTerminal, err := utils.IsTerminal(); err != nil || !isTerminal {
			// If we're not in a terminal, we can't use the dashboard
			return
		}
		dashboardOutput(formattedDestination)
	default:
		err = textOutput(formattedDestination)
	}
	if err != nil {
		fmt.Println("Flux Capacitor Deactivated:", err)
		return
	}
}

func jsonOutput(destination string) error {
	dashboard := Dashboard{
		Destination:   destination,
		Present:       Now,
		LastDeparture: LastDeparture,
		Completed:     timeWarp(),
	}
	output, err := json.MarshalIndent(dashboard, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling to JSON: %w", err)
	}
	fmt.Println(string(output))
	return nil
}

func yamlOutput(destination string) error {
	dashboard := Dashboard{
		Destination:   destination,
		Present:       Now,
		LastDeparture: LastDeparture,
		Completed:     timeWarp(),
	}

	output, err := goyaml.Marshal(dashboard)
	if err != nil {
		return fmt.Errorf("error marshaling to YAML: %w", err)
	}
	fmt.Println(string(output))
	return nil
}

func textOutput(destination string) error {
	dashboard := Dashboard{
		Destination:   destination,
		Present:       Now,
		LastDeparture: LastDeparture,
		Completed:     timeWarp(),
	}

	output := fmt.Sprintf(
		"Destination:    %s\nPresent:        %s\nLast Departure: %s\nCompleted:      %v",
		dashboard.Destination,
		dashboard.Present,
		dashboard.LastDeparture,
		dashboard.Completed,
	)
	fmt.Println(output)
	return nil
}

func dashboardOutput(destination string) {
	if dashboard.New(destination, Now, LastDeparture) {
		speedometer.New()
		utils.Print("Time warp completed!  Welcome to " + destination + "!")
		return
	}
	utils.Print("Flux Capacitor Deactivated: Stay in the present, McFly!")
}

func timeWarp() bool {
	// time travel magic
	return true
}

func init() {
	toCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|text|dashboard")
}
