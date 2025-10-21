package list

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
)

var (
	devicesClient = api.NewDevicesClient()
	devicesCmd = &cobra.Command{
		Use: "devices",
		Short: "list available spotify devices",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Debug("listing available devices")

			deviceList, err := devicesClient.AvailableDevices()
			if err != nil {
				log.Error("device list api error", "error", err)
				return err
			}
	
			fmt.Printf("Available devices:\n\n")
			for _, device := range deviceList.Devcies {
				fmt.Printf("ID: %s\n", device.ID)
				fmt.Printf("Name: %s\n", device.Name)
				fmt.Printf("Volume: %d\n\n", device.Volume)
			}

			return nil
		},	
	}
)

