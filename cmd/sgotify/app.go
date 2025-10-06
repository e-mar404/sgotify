package main

import (
	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/internal/config"
)

type App struct {
	cfg config.Config
}

func NewApp() App {
	log.Info("Creating new app")
	
	// TODO: should a func NewConfig() be created?
	return App{
		cfg: config.Config{
			RedirectURI: "http://127.0.0.1/callback",
			AuthURL: "",
			TokenURL: "",
		},
	}
}

func (a App) RunCmd(cmd cmd) error {
	log.Info("Running cmd", "cmd", cmd)
	return cmd.callback(a.cfg)	
}
