package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "display the vehicle's configuration",
	Long:  `Display the configuration of the vehicle's hardware and software.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		state, err := vehicle.VehicleConfig()

		if err != nil {
			fmt.Printf("Error retrieving vehicle configuration: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("CanAcceptNavigationRequests: %s\n", strconv.FormatBool(state.CanAcceptNavigationRequests))
		fmt.Printf("CanActuateTrunks: %s\n", strconv.FormatBool(state.CanActuateTrunks))
		fmt.Printf("CarSpecialType: %s\n", state.CarSpecialType)
		fmt.Printf("CarType: %s\n", state.CarType)
		fmt.Printf("ChargePortType: %s\n", state.ChargePortType)
		fmt.Printf("EuVehicle: %s\n", strconv.FormatBool(state.EuVehicle))
		fmt.Printf("ExteriorColor: %s\n", state.ExteriorColor)
		fmt.Printf("HasAirSuspension: %s\n", strconv.FormatBool(state.HasAirSuspension))
		fmt.Printf("HasLudicrousMode: %s\n", strconv.FormatBool(state.HasLudicrousMode))
		fmt.Printf("KeyVersion: %d\n", state.KeyVersion)
		fmt.Printf("HasMotorizedChargePort: %s\n", strconv.FormatBool(state.HasMotorizedChargePort))
		fmt.Printf("PerfConfig: %s\n", state.PerfConfig)
		fmt.Printf("PLG: %s\n", strconv.FormatBool(state.PLG))
		fmt.Printf("RearSeatHeaters: %d\n", state.RearSeatHeaters)
		fmt.Printf("RearSeatType: %d\n", state.RearSeatType)
		fmt.Printf("RHD: %s\n", strconv.FormatBool(state.RHD))
		fmt.Printf("RoofColor: %s\n", state.RoofColor)
		fmt.Printf("SeatType: %d\n", state.SeatType)
		fmt.Printf("SpoilerType: %s\n", state.SpoilerType)
		fmt.Printf("SunRoofInstalled: %d\n", state.SunRoofInstalled)
		fmt.Printf("ThirdRowSeats: %s\n", state.ThirdRowSeats)
		fmt.Printf("Timestamp: %d\n", state.Timestamp)
		fmt.Printf("TrimBadging: %s\n", state.TrimBadging)
		fmt.Printf("WheelType: %s\n", state.WheelType)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
