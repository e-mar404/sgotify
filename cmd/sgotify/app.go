package main

import (
	"log/slog"
	"net/http"
)

type App struct {
	// TODO: create client
	// client *sgotifyapi.NewClient()
	cfg Config
}

func (a App) Start() {
	slog.Info("started app with config")
	// this can also be done with a command like `sgotify login`
	slog.Info("no auth token found starting auth process")
	
	a.Authenticate()
}

func (a App) Authenticate() {
	router := newAuthRouter(a.cfg)

	slog.Info("listening on port :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		slog.Error("%v", err)
	}
}

