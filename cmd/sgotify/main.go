package main

import (
	"fmt"
	"os"

	command "github.com/e-mar404/sgotify/internal/commands"
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

	cmd := command.Cmd {
		Name: args[0],
		Args: args[1:],
	}

	if err := command.List.Run(cmd); err != nil {
		fmt.Printf("error running cmd: %s, err: %v\n", cmd.Name, err)
	}
}
