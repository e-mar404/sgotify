package cmd

import "github.com/charmbracelet/log"

func helpHandler(_ command) error {
	log.Info("printing out help...")
	return nil
}

func init() {
	availableCommands["help"] = helpHandler
}
