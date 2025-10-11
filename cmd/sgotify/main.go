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
		// if there are no arguments then run the tui 
		if err := tui.Run(); err != nil {
			fmt.Printf("error running tui: %v\n", err)
		}
		return
	}

	// if an argument is passed run that argument
	cmd := command.Cmd {
		Name: args[0],
		Args: args[1:],
	}

	if err := command.List.Run(cmd); err != nil {
		fmt.Printf("error running cmd: %s, err: %v\n", cmd.Name, err)
	}
}
