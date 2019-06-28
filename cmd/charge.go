package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kabaka/tesla"
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

		limit, _ := cmd.Flags().GetInt("limit")

		if limit != 0 {
			if limit < 0 || limit > 100 {
				fmt.Println("Invalid charge limit: value must be between 1 and 100.")
				os.Exit(1)
			}

			err := vehicle.SetChargeLimit(limit)

			if err != nil {
				fmt.Printf("Error while setting charge limit: %s\n", err)
				os.Exit(1)
			}
		}

		err := vehicle.StartCharging()

		if err != nil {
			fmt.Printf("Error while trying to start charging: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("Charging started.")
	},
}

var chargeLimitCmd = &cobra.Command{
	Use:   "limit",
	Short: "set charge limit",
	Long:  `Set the maximum charge limit to be used in the current or next charging session.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()

		limit, err := strconv.ParseInt(args[0], 0, 0)

		if err != nil {
			fmt.Println("Invalid charge limit: value must be between 1 and 100.")
			os.Exit(1)
		}

		if limit != 0 {
			if limit < 0 || limit > 100 {
				fmt.Println("Invalid charge limit: value must be between 1 and 100.")
				os.Exit(1)
			}

			err := vehicle.SetChargeLimit(int(limit))

			if err != nil {
				fmt.Printf("Error while setting charge limit: %s\n", err)
				os.Exit(1)
			}
		}

		fmt.Printf("Charge limit set to %d.\n", limit)
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

		// TODO: format this in a more readable (not prose) way

		if state.ChargingState == "Charging" {
			fmt.Printf("The vehicle is currently charging at %.1f miles/hour (%.0f volts, %.0f/%.0f amps). %.2f hours remaining.\n",
				state.ChargeRate, state.ChargerVoltage, state.ChargerActualCurrent, state.ChargerPilotCurrent, state.TimeToFullCharge)
			fmt.Printf("%.2f miles added.\n", state.ChargeMilesAddedRated)
		} else {
			fmt.Println("The vehicle is not charging.")

			if state.ScheduledChargingPending {
				fmt.Printf("Charging is scheduled to start at %s.\n", time.Unix(int64(state.ScheduledChargingStartTime.(float64)), 0).In(time.Local).Format("15:04:05 MST"))
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

// KwhCapacity determines the max battery capacity in kWh based on option codes
func KwhCapacity(v struct{ *tesla.Vehicle }) int {
	options := strings.Split(v.OptionCodes, ",")

	// TODO this is based on the list at
	// https://tesla-api.timdorr.com/vehicle/optioncodes but that list has proven
	// to be inaccurate and incomplete. As it is _so_ incomplete that it is
	// basically unusable, this function will not be used at this time.
	for _, option := range options {
		switch option {
		case "BT40":
			return 40
		case "BR03":
		case "BT60":
			return 60
		case "BR05":
		case "BTX5":
		case "BTX7":
		case "BT37":
			return 75
		case "BT85":
		case "BTX8":
			return 85
		case "BTX4":
			return 90
		case "BTX6":
			return 100
		}
	}

	fmt.Printf("No recognized battery option codes in: %s\n", v.OptionCodes)
	os.Exit(1)

	return 0 // unreachable
}

// KwhRemaining calculates the vehicle's remaining kWh
func KwhRemaining(v struct{ *tesla.Vehicle }) float64 {
	state, _ := v.ChargeState()
	return float64(KwhCapacity(v)) / float64(state.BatteryLevel)
}

func init() {
	rootCmd.AddCommand(chargeCmd)
	chargeCmd.AddCommand(chargeStartCmd)
	chargeCmd.AddCommand(chargeStopCmd)
	chargeCmd.AddCommand(chargeStatusCmd)
	chargeCmd.AddCommand(chargeLimitCmd)

	chargeStartCmd.Flags().IntP("limit", "l", 0, "Charge limit (%)")
}
