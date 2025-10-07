package command

import (
	"fmt"

	"github.com/e-mar404/sgotify/internal/config"
)

func Help(_ *config.Config) error {
	fmt.Println("help menu")
	return nil
}

