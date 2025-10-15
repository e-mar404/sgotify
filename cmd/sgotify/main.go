package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	command "github.com/e-mar404/sgotify/internal/commands"
	"github.com/e-mar404/sgotify/internal/config"
	constants "github.com/e-mar404/sgotify/internal/const"
	"github.com/e-mar404/sgotify/internal/tui"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		if err := tui.Run(); err != nil {
			fmt.Printf("error running tui: %v\n", err)
		}
		return
	}

	cfg, err := config.Load()
	if err != nil {
		log.Error("could not load a config file", "error", err)
	}

	state := &constants.State {
		Client: &http.Client{},
		Cfg: cfg,
	}

	cmd := command.Cmd {
		Name: args[0],
		Args: args[1:],
	}

	if err := command.List.Run(state, cmd); err != nil {
		fmt.Printf("error running cmd: %s, err: %v\n", cmd.Name, err)
	}
}
