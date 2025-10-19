package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/internal/tui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Verbose bool 
	Debug bool

	rootCmd = &cobra.Command{
		Use: "sgotify",
		Short: "Spotify for the cli/tui",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("starting tui")

			if err := tui.Run(); err != nil {
				log.Error("could not run tui", "error", err)
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "will show all logs except debug")
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "will show all logs")

	rootCmd.AddCommand(loginCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	log.SetOutput(os.Stderr)
	// TODO: should expand the title on the log to have a max width of 5 on the logs that get cut off (Fatal, Debug, Error)
	switch {
		case Verbose:
			// will show all logs except debug
			log.SetLevel(log.Level(0))
		case Debug:
			// debug level is -4
			log.SetLevel(log.Level(-5))
		default:
			// highest level is 12, so this will hide all logs
			log.SetLevel(log.Level(13))
	}
	
	viper.SetDefault("client_id", "")
	viper.SetDefault("client_secret", "")
	viper.SetDefault("access_token", "")
	viper.SetDefault("refresh_token", "")
	viper.SetDefault("last_refresh", 0)
	viper.SetDefault("redirect_uri", "http://127.0.0.1:8080/callback")
	viper.SetDefault("spotify_account_url", "https://accounts.spotify.com")
	viper.SetDefault("spotify_api_url", "https://api.spotify.com/v1")

	home, err := os.UserHomeDir()
	if err != nil {
		log.Error("could not get home dir", "error", err)
		cobra.CheckErr(err)
	}
	defaultPath := filepath.Join(home, ".sgotify.json")
	viper.SetConfigFile(defaultPath)
	
	if err := viper.ReadInConfig(); err != nil {
		log.Error("could not read config file", "error", err)
		log.Info("creating new config file with default values")
		viper.SafeWriteConfigAs(defaultPath)
		viper.SetConfigFile(defaultPath)
	}

	log.Info("config file loaded", "path", viper.ConfigFileUsed())
}

func requireAuth(cmd *cobra.Command, args []string) {
	// need to see if there is an access_token saved
	// if there isnt then prompt to run sgotify login

	// if there is an access token saved then check when was the last time it was refreshed
	// if it is stale then get another access_token with api.RefreshAccessToken()
	log.Info("checking to see if user is logged in + has a valid access_token")
}
