package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// softwareCmd represents the software command
var softwareCmd = &cobra.Command{
	Use:   "software",
	Short: "manage the vehicle's software",
	Long:  `Manage the vehicle's software.`,
}

var softwareScheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "schedule software update",
	Long: `Schedule the installation of pending software updates.

This cannot update the vehicle's firmware unless an update is ready.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()

		delay, err := cmd.Flags().GetInt64("delay")

		if err != nil {
			fmt.Printf("Error getting delay: %s\n", err)
			os.Exit(1)
		}

		if delay < 0 {
			fmt.Println("Invalid delay: value must be 0 or above.")
			os.Exit(1)
		}

		err = vehicle.ScheduleSoftwareUpdate(delay)

		if err != nil {
			fmt.Printf("Error scheduling update: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Software update scheduled to start in %d seconds.\n", delay)
	},
}

var softwareCancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "cancel pending software update",
	Long: `Cancel the pending software update.

If a software update is scheduled and the delay countdown has not elapsed, the update may be canceled.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		err := vehicle.CancelSoftwareUpdate()

		if err != nil {
			fmt.Printf("Error canceling software update: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("Software update canceled.")
	},
}

var softwareVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "get current version and upgrade availability",
	Long: `Determine what software version is running on the vehicle.

If a software update is available, display that here as well.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		state, err := vehicle.VehicleState()

		if err != nil {
			fmt.Printf("Error retrieving vehicle state: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Version: %s\n", state.CarVersion)

		if state.SoftwareUpdate.Status != "" {
			fmt.Printf("A software update is %s.\n", state.SoftwareUpdate.Status)
			fmt.Printf("Expected Installation Time: %d seconds\n", state.SoftwareUpdate.ExpectedDurationSec)
		}
	},
}

func init() {
	rootCmd.AddCommand(softwareCmd)
	softwareCmd.AddCommand(softwareScheduleCmd)
	softwareCmd.AddCommand(softwareCancelCmd)
	softwareCmd.AddCommand(softwareVersionCmd)

	softwareScheduleCmd.Flags().Int64P("delay", "d", 1200, "time in seconds to wait before installing update")
}
