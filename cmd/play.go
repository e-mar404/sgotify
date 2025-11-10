package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	context  string
	uris     []string
	offset   int
	position int

	playCmd = &cobra.Command{
		Use:    "play",
		Short:  "will start/resume playback on the set player",
		PreRun: checkSavedPlayer,
		Run: func(cmd *cobra.Command, args []string) {
			playerArgs := api.PlayerArgs{
				BaseURL:     viper.GetString("spotify_api_url"),
				AccessToken: viper.GetString("access_token"),
				DeviceID:    viper.GetString("device_id"),
				PlayRequestBody: api.PlayRequestBody{
					ContextURI: context,
					URIS:       uris,
					Offset: api.Offset{
						Position: offset,
					},
					PositionMs: position,
				},
			}

			var reply api.PlayerReply
			err := client.Call("Player.Play", &playerArgs, &reply)
			if err != nil {
				fmt.Printf("could not play media on device\n")
				log.Fatal("unable to play media", "device", playerArgs.DeviceID, "error", err)
			}

			if reply.Error.Status != 0 {
				fmt.Printf("could not play media on device\n")
				log.Fatal("unsuccessful action from spotify api", "status", reply.Error.Status, "message", reply.Error.Message)
			}

			fmt.Printf("Playback resumed on device\n")
		},
	}
)

func init() {
	playCmd.Flags().StringVar(&context, "context", "", "context spotify uri for play request")
	playCmd.Flags().StringSliceVar(&uris, "uris", []string{}, "list of spotify uris to play")
	playCmd.Flags().IntVar(&offset, "offset", 0, "offset detailing where in the context should playing start")
	playCmd.Flags().IntVar(&position, "position", 0, "position in ms to start playing at")
	rootCmd.AddCommand(playCmd)
}
