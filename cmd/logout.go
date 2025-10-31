package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	logoutCmd = &cobra.Command {
		Use: "logout",
		Short: "will remove any identifiable information from configuration",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("logging out...\n")
			
			viper.Set("client_id", "")
			viper.Set("client_secret", "")
			viper.Set("last_refresh", 0)
			viper.Set("access_token", "")
			viper.Set("refresh_token", "")

			if err := viper.WriteConfig(); err != nil {
				fmt.Printf("something unexpected happened while trying to log out :(\n")
				log.Fatal("could not write to config", "error", err)
			}

			fmt.Printf("successfully logged out!\n")
		},
	}
)

func init() {
	rootCmd.AddCommand(logoutCmd)
}
