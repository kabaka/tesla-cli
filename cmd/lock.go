package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// lockCmd represents the lock command
var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "control vehicle locks",
	Long:  `Lock or unlock vehicle doors.`,
}

var lockUnlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "unlock doors",
	Long:  `Unlock the vehicle's doors.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		err := vehicle.UnlockDoors()

		if err != nil {
			fmt.Printf("Error while unlocking doors: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("Doors unlocked.")
	},
}

var lockLockCmd = &cobra.Command{
	Use:   "lock",
	Short: "lock doors",
	Long:  `Lock the vehicle's doors.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		err := vehicle.LockDoors()

		if err != nil {
			fmt.Printf("Error while locking doors: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("Doors locked.")
	},
}

var lockStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "get door lock status",
	Long:  `Determine whether or not the doors are locked.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()
		state, err := vehicle.VehicleState()

		if err != nil {
			fmt.Printf("Error retrieving lock state: %s\n", err)
			os.Exit(1)
		}

		if state.Locked {
			fmt.Println("The doors are locked.")
		} else {
			fmt.Println("The doors are unlocked.")
		}
	},
}

func init() {
	rootCmd.AddCommand(lockCmd)
	lockCmd.AddCommand(lockUnlockCmd)
	lockCmd.AddCommand(lockLockCmd)
	lockCmd.AddCommand(lockStatusCmd)
}
