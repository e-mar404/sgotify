package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	constants "github.com/e-mar404/sgotify/internal/const"
)

func Run() error {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal("error setting log to file", err)
		return err
	}
	defer f.Close()

	m := newHomeUI()
	constants.P = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := constants.P.Run(); err != nil {
		log.Fatal("error running program", err)
		return err
	}

	return nil
}
