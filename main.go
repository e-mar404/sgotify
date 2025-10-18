package main

import (
	"github.com/e-mar404/sgotify/cmd"
)

func main() {
	cmd.Execute()
	// args := os.Args[1:]
	//
	// if len(args) == 0 {
	// 	if err := tui.Run(); err != nil {
	// 		fmt.Printf("error running tui: %v\n", err)
	// 	}
	// 	return
	// }
	//
	// cfg, err := config.Load()
	// if err != nil {
	// 	log.Error("could not load a config file", "error", err)
	// }
	//
	// state := &constants.State {
	// 	Client: &http.Client{},
	// 	Cfg: cfg,
	// }
	//
	// cmd := command.Cmd {
	// 	Name: args[0],
	// 	Args: args[1:],
	// }
	//
	// if err := command.List.Run(state, cmd); err != nil {
	// 	fmt.Printf("error running cmd: %s, err: %v\n", cmd.Name, err)
	// }
}
