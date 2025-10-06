package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/internal/config"
)

func main() {
	if len(os.Args) <= 1 {
		log.Info("No arguments were provided")
		cmdList["help"].callback(config.Config{})
		os.Exit(1)
	}

	rawCmd := os.Args[1]

	cmd, found := cmdList[rawCmd]
	if !found {
		log.Fatal("command not found", "cmd", rawCmd)
		os.Exit(1)
	}

	app := NewApp()
	app.RunCmd(cmd)
}
