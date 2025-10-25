package cmd

import "github.com/charmbracelet/log"

func helpHandler(cmd command) error {
	log.Info("printing out help...")
	return nil
}

func init() {
	availableCommands["help"] = serverHandler
}
