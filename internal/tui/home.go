package tui

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
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
		case key.Matches(msg, constants.Keymap.Help):
			h.help.ShowAll = !h.help.ShowAll
		}
	}

	return h, nil
}

func (h HomeUI) View() string {
	if !h.loaded {
		return "loading..."
	}

	borderOffset := 2
	helpView := h.help.View(constants.Keymap)

	// +1 accounts for the newline between content and helpView
	helpOffset := strings.Count(helpView, "\n") + 1 

	content := constants.WindowStyle.
		Height(constants.WindowSize.Height - borderOffset - helpOffset). 
		Width(constants.WindowSize.Width - borderOffset). 
		Render("login page\nyou're not logged in")

	return content + "\n" + helpView
}

func newHomeUI() HomeUI {
	log.Println("Starting new home ui")
	return HomeUI{
		help: help.New(),
	}
}
