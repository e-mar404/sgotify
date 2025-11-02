package cmd

import (
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
)

var (
	serverCmd = &cobra.Command{
		Use:    "server",
		Short:  "start rpc server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return api.StartRPCServer()
		},
	}
)

func init() {
	rootCmd.AddCommand(serverCmd)
}
