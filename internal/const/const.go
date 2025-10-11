package constants

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	P *tea.Program
	WindowSize tea.WindowSizeMsg
)

const (
	KeyringService = "sgotify-auth"	
)

var WindowStyle = lipgloss.NewStyle(). 
	Align(lipgloss.Center, lipgloss.Center).
	Border(lipgloss.RoundedBorder())

type keymap struct {
	Quit   key.Binding
	Help key.Binding
}

// Keymap reusable key mappings shared across models
var Keymap = keymap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help menu"),
	),
}

func (k keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{},
		{k.Quit, k.Help},
	}
}
