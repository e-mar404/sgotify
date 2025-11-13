package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	pauseCmd = &cobra.Command{
		Use:   "pause",
		Short: "Pause playback on current active device",
		Run: func(cmd *cobra.Command, args []string) {
			var reply api.PlayerReply
			pauseArgs := api.PlayerArgs{
				AccessToken: viper.GetString("access_token"),
				DeviceID:    viper.GetString("device_id"),
			}

			err := client.Call("Player.Pause", &pauseArgs, &reply)
			if err != nil {
				fmt.Printf("unexpected error while pausing playback")
				log.Fatal("unexpected action error", "error", err)
			}

			if reply.Error.Status != 0 {
				fmt.Printf("unable to complete request")
				log.Error("unsuccessful action", "status", reply.Error.Status, "message", reply.Error.Message)
				return
			}

			fmt.Printf("Playback paused on device\n")
		},
	}
)

func init() {
	rootCmd.AddCommand(pauseCmd)
}
