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

		if state.ChargingState == "Charging" {
			fmt.Printf("The vehicle is currently charging at %.1f miles/hour (%.0f volts, %.0f/%.0f amps). %.2f hours remaining.\n",
				state.ChargeRate, state.ChargerVoltage, state.ChargerActualCurrent, state.ChargerPilotCurrent, state.TimeToFullCharge)
			fmt.Printf("%.2f miles added.\n", state.ChargeMilesAddedRated)
		} else {
			fmt.Println("The vehicle is not charging.")

			if state.ManagedChargingActive {
				fmt.Println("Charging is scheduled to start at ...") // XXX
			} else {
				if state.ChargePortDoorOpen {
					fmt.Println("The charge port is open.")
				} else {
					fmt.Println("The charge port is closed.")
				}
			}
		}

		fmt.Printf("Limit: %d%%\n", state.ChargeLimitSoc)
		fmt.Printf("Battery Level: %d%%\n", state.BatteryLevel)
	},
}

func init() {
	rootCmd.AddCommand(chargeCmd)
	chargeCmd.AddCommand(chargeStartCmd)
	chargeCmd.AddCommand(chargeStopCmd)
	chargeCmd.AddCommand(chargeStatusCmd)
}
