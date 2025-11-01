package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	list bool

	playerService = api.NewPlayerService()

	playerCmd = &cobra.Command{
		Use: "player",
		Short: "command to interact with a spotify player state",
		PreRun: requireAuth,
		Run: func(cmd *cobra.Command, args []string) {
			switch {
			case list:
				args := api.PlayerArgs{
					BaseUrl: viper.GetString("spotify_api_url"),
					AccessToken: viper.GetString("access_token"),
				}
				reply := api.AvailableDevicesReply{}
				err := playerService.AvailableDevices(&args, &reply)
				if err != nil {
					fmt.Printf("could not get available devices\n")
					log.Fatal("could not get available devices", "error", err)
				}

				if len(reply.Devices) == 0 {
					fmt.Printf("There are no available devices. Please start spotify on one of your devices\n")
					return
				}

				fmt.Printf("Available Devices:\n")
				for _, device := range reply.Devices {
					fmt.Printf("========================\n")
					fmt.Printf("ID: %s\n", device.ID)
					fmt.Printf("Name: %s\n", device.Name)
				}
				fmt.Printf("========================\n")

			default:
				fmt.Printf("device help menu\n")
			}
		},
	}
)

func init() {
	playerCmd.PersistentFlags().BoolVarP(&list, "list", "l", false, "will list all available devices")
	rootCmd.AddCommand(playerCmd)
}
