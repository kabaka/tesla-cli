package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sentryCmd represents the sentry command
var sentryCmd = &cobra.Command{
	Use:   "sentry status|on|off",
	Short: "enable or disable Sentry Mode",
	Long:  `Enable or disable Sentry Mode.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

var sentryOnCmd = &cobra.Command{
	Use:   "on",
	Short: "enable Sentry Mode",
	Long:  `Enable Sentry Mode.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		err := vehicle.SentryModeEnable()

		if err != nil {
			fmt.Printf("Error enabling sentry mode: %s\n", err)
		}

		fmt.Println("Sentry mode enabled.")
	},
}

var sentryOffCmd = &cobra.Command{
	Use:   "off",
	Short: "disable Sentry Mode",
	Long:  `Disable Sentry Mode.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		err := vehicle.SentryModeDisable()

		if err != nil {
			fmt.Printf("Error enabling sentry mode: %s\n", err)
		}

		fmt.Println("Sentry mode disabled.")
	},
}

var sentryStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "get Sentry Mode status",
	Long:  `Determine whether Sentry Mode is currently enabled or disabled.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		state, err := vehicle.VehicleState()

		if err != nil {
			fmt.Printf("Error enabling sentry mode: %s\n", err)
		}

		if state.SentryMode {
			fmt.Println("Sentry mode is enabled.")
		} else {
			fmt.Println("Sentry mode is disabled.")
		}
	},
}

func init() {
	rootCmd.AddCommand(sentryCmd)
	sentryCmd.AddCommand(sentryOnCmd)
	sentryCmd.AddCommand(sentryOffCmd)
	sentryCmd.AddCommand(sentryStatusCmd)
}
