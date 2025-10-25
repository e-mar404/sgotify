package cmd

import "github.com/charmbracelet/log"

func serverHandler(cmd command) error {
	log.Info("starting server...")
	return nil
}

func init() {
	availableCommands["server"] = serverHandler
}
