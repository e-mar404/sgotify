package cmd

import (
	"fmt"
	"net/rpc"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	verbose bool
	client *rpc.Client

	rootCmd = cobra.Command{
		Use:   "sgotify",
		Short: "start tui client",
		PersistentPostRun: prepLogs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("tui client coming soon...\n")

		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "set verbose output")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Error("unable to run sgotify", "error", err)
		os.Exit(1)
	}
}

func initConfig() {
	log.SetOutput(os.Stderr)
	// TODO: should expand the title on the log to have a max width of 5 on the logs that get cut off (Fatal, Debug, Error)

	viper.SetDefault("spotify_account_url", "https://accounts.spotify.com")
	viper.SetDefault("spotify_api_url", "https://api.spotify.com/v1")
	viper.SetDefault("redirect_uri", "http://127.0.0.1:8080/callback")
	viper.SetDefault("client_id", "")
	viper.SetDefault("client_secret", "")
	viper.SetDefault("last_refresh", 0)
	viper.SetDefault("access_token", "")
	viper.SetDefault("refresh_token", "")

	home, err := os.UserHomeDir()
	if err != nil {
		log.Error("unable to find user's home dir", "error", err)
		os.Exit(1)
	}

	configPath := filepath.Join(home, ".config", "sgotify", "config.json")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Info("no pre-existent config file, creating default", "path", configPath)
		os.MkdirAll(filepath.Dir(configPath), 0766)
		viper.SafeWriteConfigAs(configPath)
		viper.SetConfigFile(configPath)
	}
}
