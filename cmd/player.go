package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	list   bool
	device string

	playerCmd = &cobra.Command{
		Use:   "player",
		Short: "command to interact with a spotify player state",
		Run: func(cmd *cobra.Command, _ []string) {
			args := &api.PlayerArgs{
				BaseURL:     viper.GetString("spotify_api_url"),
				AccessToken: viper.GetString("access_token"),
			}
			reply := &api.AvailableDevicesReply{}
			err := client.Call("Player.AvailableDevices", args, reply)
			if err != nil {
				fmt.Printf("could not get available devices\n")
				log.Fatal("could not get available devices", "error", err)
			}

			if len(reply.Devices) == 0 {
				fmt.Printf("There are no available devices. Please start spotify on one of your devices\n")
				return
			}

			switch {
			case list:
				printDeviceList(reply)
			case device == "":
				deviceID := viper.GetString("device_id")
				if deviceID == "" {
					fmt.Printf("No device set in configuration. Please run sgotify player --set-device=<device_id> to set a device\n")
					log.Info("no device set in configuration")
				}

				for _, d := range reply.Devices {
					if d.ID == deviceID {
						fmt.Printf("Current selected device:\n\n")
						fmt.Printf("ID: %s\n", d.ID)
						fmt.Printf("Name: %s\n", d.Name)

						break
					}
				}

			case device != "":
				for _, d := range reply.Devices {
					if d.ID == device {
						viper.Set("device_id", device)
						if err := viper.WriteConfig(); err != nil {
							fmt.Printf("unable to write device id to config\n")
							log.Fatal("could not write to config", "error", err)
						}
						fmt.Printf("Player %s set successfully!\n", d.Name)
						return
					}
				}

				fmt.Printf("given device id was not a valid device id from the list of available devices\n\n")
				log.Info("device given not found", "device", device)
				printDeviceList(reply)
			}
		},
	}
)

func init() {
	playerCmd.Flags().BoolVarP(&list, "list-devices", "l", false, "will list all available devices")
	playerCmd.Flags().StringVarP(&device, "set-device", "s", "", "will set given device id as the player")
	playerCmd.MarkFlagsMutuallyExclusive("list-devices", "set-device")

	rootCmd.AddCommand(playerCmd)
}

func printDeviceList(reply *api.AvailableDevicesReply) {
	fmt.Printf("Available Devices:\n")
	for _, device := range reply.Devices {
		fmt.Printf("========================\n")
		fmt.Printf("ID: %s\n", device.ID)
		fmt.Printf("Name: %s\n", device.Name)
	}
	fmt.Printf("========================\n")
}

func checkSavedPlayer(_ *cobra.Command, _ []string) {
	log.Info("checking to see if there is a saved player\n")
	deviceID := viper.GetString("device_id")
	if deviceID == "" {
		fmt.Printf("no device saved in config. Please run sgotify player --set-device <device_id>\n")
		log.Fatal("no device saved in config")
	}
}
