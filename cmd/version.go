package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Version of timetravel CLI",
	Example: `timetravel version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}