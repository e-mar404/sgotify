package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	playCmd = &cobra.Command{
		Use:    "play",
		Short:  "will start/resume playback on the set player",
		PreRun: checkSavedPlayer,
		Run: func(cmd *cobra.Command, args []string) {
			playerArgs := api.PlayerArgs{
				BaseURL:     viper.GetString("spotify_api_url"),
				AccessToken: viper.GetString("access_token"),
				DeviceID:    viper.GetString("device_id"),
			}

			var reply api.PlayReply
			err := client.Call("Player.Play", &playerArgs, &reply)
			if err != nil {
				fmt.Printf("could not play media on device\n")
				log.Fatal("unable to play media", "device", playerArgs.DeviceID, "error", err)
			}

			if reply.Error.Status != 0 {
				fmt.Printf("could not play media on device\n")
				log.Fatal("unsuccessful action from spotify api", "status", reply.Error.Status, "message", reply.Error.Message)
			}

			fmt.Printf("Resuming media on device\n")
		},
	}
)

func init() {
	rootCmd.AddCommand(playCmd)
}
