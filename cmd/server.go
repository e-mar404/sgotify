package cmd

import (
	"github.com/e-mar404/sgotify/api"
)

func serverHandler(_ command) error {
	return api.StartRPCServer()
}

func init() {
	availableCommands["server"] = serverHandler
}
