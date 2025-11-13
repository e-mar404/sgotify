package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	prevCmd = &cobra.Command{
		Use:   "prev",
		Short: "go to the previous track",
		Run: func(cmd *cobra.Command, args []string) {
			var reply api.PlayerReply
			prevArgs := api.PlayerArgs{
				AccessToken: viper.GetString("access_token"),
				DeviceID:    viper.GetString("device_id"),
			}

			err := client.Call("Player.Prev", &prevArgs, &reply)
			if err != nil {
				fmt.Printf("unexpected error occurred\n")
				log.Fatal("unexpected error", "error", err)
			}

			if reply.Error.Status != 0 {
				fmt.Printf("unsuccessful spotify action\n")
				log.Fatal("unsuccessful action", "status", reply.Error.Status, "message", reply.Error.Message)
			}

			fmt.Printf("Going to previous track from the queue\n")
		},
	}
)

func init() {
	rootCmd.AddCommand(prevCmd)
}
