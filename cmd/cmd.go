package cmd

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/internal/tui"
	"github.com/spf13/viper"
)

type command struct {
	name      string
	arguments []string
}

type commands map[string]func(command) error

var availableCommands = commands{}

func (cmds commands) AddCommand(name string, run func(command) error) {
	cmds[name] = run
}

func initConfig() error {
	viper.SetDefault("spotify_account_url", "https://accounts.spotify.com")
	viper.SetDefault("spotify_api_url", "https://api.spotify.com/v1")
	viper.SetDefault("redirect_uri", "http://127.0.0.1:8080")
	viper.SetDefault("client_id", "")
	viper.SetDefault("client_secret", "")
	viper.SetDefault("last_refresh", 0)
	viper.SetDefault("access_token", "")
	viper.SetDefault("refresh_token", "")

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(home, ".config", "sgotify", "config.json")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Info("no pre-existent config file, creating default", "path", configPath)
		os.MkdirAll(filepath.Dir(configPath), 0766)
		viper.SafeWriteConfigAs(configPath)
		viper.SetConfigFile(configPath)
	}

	return nil
}

func Execute() {
	if err := initConfig(); err != nil {
		log.Error("unable to initialize config", "error", err)
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		log.Info("no cmd starting tui...")
		if err := tui.Run(); err != nil {
			log.Error("something unexpected happened while running the tui", "error", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	cmdName := os.Args[1]
	args := os.Args[2:]

	run, ok := availableCommands[cmdName]
	if !ok {
		log.Error("cmd not found", "name", cmdName)
		os.Exit(1)
	}

	cmd := command{
		cmdName,
		args,
	}

	if err := run(cmd); err != nil {
		log.Error("command failed", "error", err)
		os.Exit(1)
	}
}
