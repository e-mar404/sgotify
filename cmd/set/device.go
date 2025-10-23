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
	id string
	deviceCmd     = &cobra.Command{
		Use:   "device",
		Short: "Will set the device passed to be the output device for spotify",
		RunE: func(cmd *cobra.Command, args []string) error {
			deviceList, err := devicesClient.AvailableDevices()
			if err != nil {
				log.Error("device list api error", "error", err)
				return err
			}

			var defaultDevice api.DeviceItem
			for _, device := range deviceList.Devcies {
				if device.ID == id {
					defaultDevice = device
				}
			}

			if defaultDevice.Name == "" {
				log.Error("not a valid device", "id", id)
				return fmt.Errorf("Could not find available device with id of %s", id)
			}

			viper.Set("default_device_id", defaultDevice.ID)
			viper.Set("default_device_name", defaultDevice.Name)
			if err := viper.WriteConfig(); err != nil {
				log.Error("could not save default device to config", "error", err)
				return fmt.Errorf("unable to save device with id %s as default", id)
			}
			
			fmt.Printf("successfully set device %s (%s)\n", id, defaultDevice.Name)

			return nil
		},
	}
)

func init() {
	deviceCmd.PersistentFlags().StringVar(&id, "id", "", "device id to set as default")
	deviceCmd.MarkFlagRequired("id")
}
