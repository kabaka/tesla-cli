package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// climateCmd represents the climate command
var climateCmd = &cobra.Command{
	Use:   "climate",
	Short: "control the car's climate control system",
	Long:  `Query and adjust heating, cooling, and ventilation settings.`,
}

var climateOnCmd = &cobra.Command{
	Use:   "on",
	Short: "enable climate control",
	Long: `Enable climate control, which will begin heating or cooling the
    interior of the vehicle as needed.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		err := vehicle.StartAirConditioning()

		if err != nil {
			fmt.Printf("Error enabling climate control: %s\n", err)
		}

		fmt.Println("Climate control enabled.")
	},
}

var climateOffCmd = &cobra.Command{
	Use:   "off",
	Short: "disable climat econtrol",
	Long: `Disable climate control, which will stop heating or cooling of the
		interior of the vehicle.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		err := vehicle.StopAirConditioning()

		if err != nil {
			fmt.Printf("Error enabling climate control: %s\n", err)
		}

		fmt.Println("Climate control disabled.")
	},
}

var climateStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "get climate control status",
	Long: `Determine whether climate control is currently enabled or disabled,
		as well as the target temperature range.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		state, err := vehicle.ClimateState()

		if err != nil {
			fmt.Printf("Error retrieving climate control status: %s\n", err)
		}

		if state.IsClimateOn {
			fmt.Println("Climate control is enabled.")
		} else {
			fmt.Println("Climate control is disabled.")
		}
	},
}

func init() {
	rootCmd.AddCommand(climateCmd)
	climateCmd.AddCommand(climateOnCmd)
	climateCmd.AddCommand(climateOffCmd)
	climateCmd.AddCommand(climateStatusCmd)
}
