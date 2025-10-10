package tui

import (
	"log"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/e-mar404/sgotify/internal/tui/constants"
)

type HomeUI struct {
	loaded bool
	help help.Model
}

func (h HomeUI) Init() tea.Cmd {
	return nil
}

func (h HomeUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg
		h.loaded = true

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Quit):
			return h, tea.Quit
		}
	}

	return h, nil
}

func (h HomeUI) View() string {
	if !h.loaded {
		return "loading..."
	}

	content := lipgloss.Place(
		constants.WindowSize.Width,
		constants.WindowSize.Height,	
		lipgloss.Center,
		lipgloss.Center,
		"login page",
	)
	helpView := lipgloss.Place(
		constants.WindowSize.Width,
		constants.WindowSize.Height,	
		lipgloss.Center,
		lipgloss.Bottom,
		h.help.View(constants.Keymap),
	)


	return content + helpView 
}

func newHomeUI() HomeUI {
	log.Println("Starting new home ui")
	return HomeUI{
		help: help.New(),
	}
}
