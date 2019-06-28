package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// honkCmd honks the vehicle's horn
var honkCmd = &cobra.Command{
	Use:   "honk",
	Short: "honk the horn",
	Long:  `Honk the vehicle's horn.`,
	Run: func(cmd *cobra.Command, args []string) {
		vehicle := GetTeslaVehicle()

		err := vehicle.HonkHorn()

		if err != nil {
			fmt.Printf("Error while unlocking doors: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("Horn successfully honked.")
	},
}

func init() {
	rootCmd.AddCommand(honkCmd)
}
