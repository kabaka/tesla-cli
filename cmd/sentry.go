package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// sentryCmd represents the sentry command
var sentryCmd = &cobra.Command{
	Use:   "sentry status|on|off",
	Short: "enable or disable Sentry Mode",
	Long:  `Enable or disable Sentry Mode.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("usage: sentry status|on|off")
			os.Exit(1)
		}

		if args[0] == "on" {
			vehicle := GetTeslaVehicle()
			err := vehicle.SentryModeEnable()

			if err != nil {
				fmt.Printf("Error enabling sentry mode: %s\n", err)
			}

			fmt.Println("Sentry mode enabled.")
			return

		} else if args[0] == "off" {
			vehicle := GetTeslaVehicle()
			err := vehicle.SentryModeDisable()

			if err != nil {
				fmt.Printf("Error enabling sentry mode: %s\n", err)
			}

			fmt.Println("Sentry mode disabled.")
			return
		}

		fmt.Println("usage: sentry status|on|off")
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(sentryCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sentryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sentryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
