package main

import (
	"github.com/charmbracelet/log"
	command "github.com/e-mar404/sgotify/internal/commands"
	"github.com/e-mar404/sgotify/internal/config"
)

type App struct {
	cfg *config.Config
}

func NewApp() App {
	log.Info("Creating new app")
	
	// TODO: should a func NewConfig() be created?
	return App{
		cfg: &config.Config{
			RedirectURI: "http://127.0.0.1/callback",
			AuthURL: "https://accounts.spotify.com/authorize",
			TokenURL: "https://accounts.spotify.com/api/token",
		},
	}
}

func (a App) RunCmd(cmd command.Cmd) error {
	log.Info("Running cmd", "cmd", cmd.Name)
	return cmd.Callback(a.cfg)	
}
