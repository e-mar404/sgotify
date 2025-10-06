package main

import (
	command "github.com/e-mar404/sgotify/internal/commands"
	"github.com/e-mar404/sgotify/internal/config"
)

type cmd struct {
	 name string
	 description string
	 callback func(config.Config) error
}

var cmdList = map[string]cmd {
	"login": {
		name: "login",
		description: "Start spotify authentication flow",
		callback: command.Login,
	},
	"help": {
		name: "help",
		description: "List all available commands",
		callback: command.Help,
	},
}
