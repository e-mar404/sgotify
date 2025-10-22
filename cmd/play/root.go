package play

import (
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
)

var (
	player  = api.NewPlayer()
	RootCmd = &cobra.Command{
		Use:   "play",
		Short: "play spotify objects",
	}
)

func init() {
	RootCmd.PersistentFlags().String("ID", "", "track id to play on default player")
}
