package cmd

import (
	"fmt"
	"os"

	"github.com/kabaka/tesla"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "tesla",
	Short: "Control Tesla vehicles.",
	Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default: $HOME/.tesla.toml)")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tesla")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

// GetTeslaVehicle returns a Tesla vehicle object with active API connection
func GetTeslaVehicle() struct{ *tesla.Vehicle } {
	// TODO: store token and try to reuse before re-authenticating
	client, err := tesla.NewClient(
		&tesla.Auth{
			ClientID:     viper.GetString("client_id"),
			ClientSecret: viper.GetString("client_secret"),
			Email:        viper.GetString("email"),
			Password:     viper.GetString("password"),
		})

	if err != nil {
		panic(err)
	}

	vehicles, err := client.Vehicles()
	if err != nil {
		panic(err)
	}
	vehicle := vehicles[0]
	_, err = vehicle.MobileEnabled()

	if err != nil {
		panic(err)
	}

	return vehicle
}
