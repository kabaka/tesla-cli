package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// lightsCmd lightss the vehicle's horn
var lightsCmd = &cobra.Command{
	Use:   "lights",
	Short: "control lights",
	Long:  `Control the vehicle's lights.`,
}

var lightsFlashCmd = &cobra.Command{
	Use:   "flash",
	Short: "flash exterior lights",
	Long:  `Flash the vehicle's exterior lights.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()

		err := vehicle.FlashLights()

		if err != nil {
			fmt.Printf("Error while flashing lights: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("Lights successfully flashed.")
	},
}

func init() {
	rootCmd.AddCommand(lightsCmd)
	lightsCmd.AddCommand(lightsFlashCmd)
}
