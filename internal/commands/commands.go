package command

import (
	"fmt"
)

func init() {
	List.Register(Cmd {
		Name: "help",
		Description: "Will print out the help menu",
		Callback: Help,
	}, 
	Cmd {
		Name: "login",
		Description: "Sign in with client id & secret to get auth & refresh tokens",
		Callback: Login,
	})
}

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

func (cmds Cmds) Register(cmdList ... Cmd) {
	for _, cmd := range cmdList {
		cmds[cmd.Name] = cmd
	}
}
