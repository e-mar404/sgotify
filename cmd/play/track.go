package play

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	trackCmd = &cobra.Command{
		Use:   "track",
		Short: "play a spotify song",
		RunE: func(cmd *cobra.Command, args []string) error {
			trackID := cmd.Flag("ID").Value.String()
			if trackID == "" {
				fmt.Println("no track id passed")
				return nil
			}

			fmt.Printf("playing track %s on device %s\n", trackID, player.DeviceName)

			return nil
		},
	}
)

func init() {
	RootCmd.AddCommand(trackCmd)
}
