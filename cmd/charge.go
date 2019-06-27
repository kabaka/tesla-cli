package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// chargeCmd represents the charge command
var chargeCmd = &cobra.Command{
	Use:   "charge",
	Short: "control vehicle charging",
	Long:  `Start or stop charing and change the charge limit.`,
}

var chargeStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start charging",
	Long:  `Start charging the vehicle if it is plugged in.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		err := vehicle.StartCharging()

		if err != nil {
			fmt.Printf("Error while trying to start charging: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("Charging started.")
	},
}

var chargeStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop charging",
	Long:  `Stop charging the vehicle.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		err := vehicle.StopCharging()

		if err != nil {
			fmt.Printf("Error while trying to stop charging: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("Charging stopped.")
	},
}

var chargeStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "get charging status",
	Long:  `Determine current charge state and limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		state, err := vehicle.ChargeState()

		if err != nil {
			fmt.Printf("Error retrieving charge state: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Charge state: %s\n", state.ChargingState)
	},
}

func init() {
	rootCmd.AddCommand(chargeCmd)
	chargeCmd.AddCommand(chargeStartCmd)
	chargeCmd.AddCommand(chargeStopCmd)
	chargeCmd.AddCommand(chargeStatusCmd)
}
