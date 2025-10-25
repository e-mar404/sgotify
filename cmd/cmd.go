package cmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/internal/tui"
)

type command struct {
	name      string
	arguments []string
}

type commands map[string]func(command) error

var availableCommands = map[string]func(command) error{}

func (cmds commands) AddCommand(name string, run func(command) error) {
	cmds[name] = run
}

func Execute() {
	if len(os.Args) == 1 {
		log.Info("no cmd starting tui...")
		if err := tui.Run(); err != nil {
			log.Error("something unexpected happened while running the tui", "error", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	cmdName := os.Args[1]
	args := os.Args[2:]

	run, ok := availableCommands[cmdName]
	if !ok {
		log.Error("cmd not found", "name", cmdName)
		os.Exit(1)
	}

	cmd := command{
		cmdName,
		args,
	}

	if err := run(cmd); err != nil {
		log.Error("command failed", "error", err)
		os.Exit(1)
	}
}
