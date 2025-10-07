package main

import (
	"fmt"
	"os"

	command "github.com/e-mar404/sgotify/internal/commands"
	"github.com/e-mar404/sgotify/internal/config"
)

func main() {
	cmds := command.List()

	if len(os.Args) <= 1 {
		fmt.Println("No arguments were provided")
		cmds["help"].Callback(&config.Config{})
		return
	}

	rawCmd := os.Args[1]
	cmd, found := cmds[rawCmd]
	if !found {
		fmt.Printf("command `%v` not found\n", rawCmd)
		cmds["help"].Callback(&config.Config{})
		return
	}

	app := NewApp()
	app.RunCmd(cmd)
}
