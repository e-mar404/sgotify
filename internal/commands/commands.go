package command

import "fmt"

type Cmd struct {
	 Name string
	 Description string
	 Args []string
	 Callback func() error
}

type Cmds map[string]Cmd

var List = Cmds {
	"help": {
		Name: "help",
		Description: "List all available commands",
		Callback: Help,
	},
}

func (l Cmds) Run(cmd Cmd) error {
	c, found := l[cmd.Name]	
	if !found {
		return fmt.Errorf("did not find cmd: %v\n", cmd.Name)
	}

	return c.Callback()
}
