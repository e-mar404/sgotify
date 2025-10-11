package command

import (
	"fmt"
)

type Cmd struct {
	 Name string
	 Description string
	 Args []string
	 Callback func() error
}

type Cmds map[string]Cmd

var List = Cmds{}

func (cmds Cmds) Run(cmd Cmd) error {
	c, found := cmds[cmd.Name]	
	if !found {
		return fmt.Errorf("did not find cmd: %v\n", cmd.Name)
	}

	return c.Callback()
}

func (cmds Cmds) Register(cmd Cmd) {
	cmds[cmd.Name] = cmd
}

func InitCmds() {
	List.Register(Cmd {
		Name: "help",
		Description: "Will print out the help menu",
		Callback: Help,
	})
}
