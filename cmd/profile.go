package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
)

var (
	profileCmd  = &cobra.Command {
		Use: "profile",
		Short: "display some stats about your spotify profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("getting profile stuff...")

			res, err := api.UserProfile()
			if err != nil {
				log.Error("could not get user profile", "error", err)
				return err
			}

			log.Info("got back a user profile", "profile", res)

			fmt.Printf("Display Name: %s\n", res.DisplayName)
			fmt.Printf("Email: %s\n", res.Email)

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(profileCmd)
}

