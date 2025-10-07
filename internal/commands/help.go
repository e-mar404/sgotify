package command

import (
	"fmt"

	"github.com/e-mar404/sgotify/internal/config"
)

func Help(_ *config.Config) error {
	fmt.Println("help menu")

	cmds := List()
	fmt.Println("============")
	for _, cmd := range cmds {
		fmt.Printf("%s\n", cmd.name)
		fmt.Printf("%s\n", cmd.description)
	}

	return nil
}

