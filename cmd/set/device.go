package set

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	devicesClient = api.NewDevicesClient()
	deviceCmd     = &cobra.Command{
		Use:   "device",
		Short: "Will set the device passed to be the output device for spotify",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Debug("listing available devices")

			deviceList, err := devicesClient.AvailableDevices()
			if err != nil {
				log.Error("device list api error", "error", err)
				return err
			}

			input := args[0]
			var defaultDevice api.DeviceItem
			for _, device := range deviceList.Devcies {
				if device.Name == input || device.ID == input {
					defaultDevice = device
				}
			}

			// if loop could not find a matching device
			if defaultDevice.Name == "" {
				log.Error("not a valid device")
				return fmt.Errorf("Could not find available device with name or id of %s", input)
			}

			viper.Set("default_device_id", defaultDevice.ID)
			viper.Set("default_device_name", defaultDevice.Name)
			if err := viper.WriteConfig(); err != nil {
				log.Error("could not save default device to config", "error", err)
				return fmt.Errorf("unable to save device %s as default", input)
			}

			return nil
		},
	}
)
