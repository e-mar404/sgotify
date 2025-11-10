package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	nextCmd = &cobra.Command{
		Use:   "next",
		Short: "skip to next track",
		Run: func(cmd *cobra.Command, args []string) {
			var reply api.PlayerReply
			nextArgs := api.PlayerArgs{
				BaseURL:     viper.GetString("spotify_api_url"),
				AccessToken: viper.GetString("access_token"),
				DeviceID:    viper.GetString("device_id"),
			}

			err := client.Call("Player.Next", &nextArgs, &reply)
			if err != nil {
				fmt.Printf("unexpected error happened\n")
				log.Fatal("unexpected error", "error", err)
			}

			if reply.Error.Status != 0 {
				fmt.Printf("unable to skip track\n")
				log.Error("unsuccessful action", "status", reply.Error.Status, "message", reply.Error.Message)
			}

			fmt.Printf("Playing next song\n")
		},
	}
)

func init() {
	rootCmd.AddCommand(nextCmd)
}
