package constants

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	P *tea.Program
	WindowSize tea.WindowSizeMsg
)

const (
	KeyringService = "sgotify-auth"	
)

type keymap struct {
	Login key.Binding
	Quit   key.Binding
}

// Keymap reusable key mappings shared across models
var Keymap = keymap{
	Login: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "login"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
}

func (k keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}

func (k keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Login},
		{k.Quit},
	}
}
