package command

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/internal/config"
	"github.com/e-mar404/sgotify/internal/models"
)

func Login(cfg *config.Config) error {
	p := tea.NewProgram(models.NewLoginModel())
	if _, err := p.Run(); err != nil {
		log.Error("failed to launch tui", "error", err)
		return err
	}

	return nil
}
