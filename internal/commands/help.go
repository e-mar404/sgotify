package command

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/e-mar404/sgotify/internal/config"
)

func Help(_ *config.Config) error {
	s := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render
	
	usage := strings.Builder{}
	usage.WriteString("Usage:\n\n")
	usage.WriteString("  sgotify " + s("[command]") + "\n")
	fmt.Println(usage.String())

	t := table.New()

	cmds := List()
	for _, cmd := range cmds {
		t.Row(cmd.Name, s(cmd.description))
	}
	fmt.Println("Command list:")
	fmt.Println(t.Render())

	return nil
}

