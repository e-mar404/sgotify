package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	logoutCmd = &cobra.Command{
		Use:   "logout",
		Short: "will delete any saved user data from the configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info("deleting any saved user data from config file")

			viper.Set("access_token", "")
			viper.Set("refresh_token", "")
			viper.Set("last_refresh", 0)
			viper.Set("default_device_name", "")
			viper.Set("default_device_id", "")
			viper.Set("client_id", "")
			viper.Set("client_secret", "")

			return viper.WriteConfig()
		},
	}
)

func init() {
	rootCmd.AddCommand(logoutCmd)
}
