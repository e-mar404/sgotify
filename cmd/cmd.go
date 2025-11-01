package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/e-mar404/sgotify/internal/tui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	verbose bool

	rootCmd = cobra.Command{
		Use:   "sgotify",
		Short: "start tui client",
		Run: func(cmd *cobra.Command, args []string) {
			if err := tui.Run(); err != nil {
				log.Error("something unexpected happened while running the tui", "error", err)
				os.Exit(1)
			}
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

func batch(cmds ...func(*cobra.Command, []string)) func(*cobra.Command, []string) {
	return  func(c *cobra.Command, s []string) {
		for _, cmd := range cmds {
			cmd(c, s)
		}
	}
}

func prepLogs(cmd *cobra.Command, args []string) {
	baseLevel := log.Level(13)
	verboseLevel := log.Level(0)
	if cmd.Use == "server" { // only the server cmd will have logs by default
		baseLevel = log.Level(0)
		verboseLevel = log.Level(-5)
	}

	if verbose {
		log.SetLevel(verboseLevel)
	} else {
		log.SetLevel(baseLevel)
	}
}

func requireAuth(cmd *cobra.Command, args []string) {
	log.Info("checking access token status")
	assert := func(condition bool) {
		if condition {
			fmt.Fprintln(os.Stderr, "Please run `sgotify login` first.")
			log.Fatal("not signed in")
		}
	}

	accessToken := viper.GetString("access_token")
	assert(accessToken == "")

	last_refresh := viper.GetInt64("last_refresh")
	if time.Now().Add(-time.Minute*55).Unix() <= last_refresh {
		log.Info("Access token is still good, not refreshing")
		return
	}

	log.Info("asking for a new access token")

	refreshArgs := api.RefreshArgs {
		RefreshToken: viper.GetString("refresh_token"),
		BaseURL: viper.GetString("spotify_api_url"),
	}
	reply := api.CredentialsReply{}
	err := authService.RefreshAccessToken(&refreshArgs, &reply)
	assert(err != nil)

	viper.Set("access_token", reply.AccessToken)
	viper.Set("last_refresh", time.Now().Unix())
	if reply.RefreshToken != "" {
		viper.Set("refresh_token", reply.RefreshToken)
	}

	err = viper.WriteConfig()
	assert(err != nil)
}
