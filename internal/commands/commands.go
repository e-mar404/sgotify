package command

import (
	"github.com/e-mar404/sgotify/internal/config"
)

type Cmd struct {
	 name string
	 description string
	 Callback func(*config.Config) error
}

func List() map[string]Cmd {
	return map[string]Cmd {
		"login": {
			name: "login",
			description: "Start spotify authentication flow",
			Callback: Login,
		},
		"help": {
			name: "help",
			description: "List all available commands",
			Callback: Help,
		},
	}
}
