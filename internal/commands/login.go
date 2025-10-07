package command

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/internal/config"
	"github.com/zalando/go-keyring"
)


const (
	cid = iota
	cs
)

type authModel struct {
	inputs []textinput.Model
	focused int
}

func newAuthModel() authModel {
	var inputs []textinput.Model = make([]textinput.Model, 3)
	inputs[cid] = textinput.New()
	inputs[cid].Placeholder = "CLIENT_ID"
	inputs[cid].Focus()
	inputs[cid].Width = 50
	inputs[cid].Prompt = ""

	inputs[cs] = textinput.New()
	inputs[cs].Placeholder = "CLIENT_SECRET"
	inputs[cs].Width = 50
	inputs[cs].Prompt = ""

	return authModel{
		inputs:  inputs,
		focused: 0,
	}
}

func (a authModel) Init() tea.Cmd {
	return textinput.Blink
}

func (a authModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(a.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if a.focused == len(a.inputs)-1 {
				return a, tea.Quit
			}
			a.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return a, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			a.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			a.nextInput()
		}
		for i := range a.inputs {
			a.inputs[i].Blur()
		}
		a.inputs[a.focused].Focus()

	// We handle errors just like any other message
	case error:
		return a, nil
	}

	for i := range a.inputs {
		a.inputs[i], cmds[i] = a.inputs[i].Update(msg)
	}
	return a, tea.Batch(cmds...)
}

func (a authModel) View() string {
	return fmt.Sprintf(`
	%s
	%s

	%s
	%s

	%s`,
	"Enter Your Spotify Client ID:",
	a.inputs[cid].View(),
	"Enter Your Spotify Client Secret",
	a.inputs[cs].View(),
	"Continue ->")
} 

// nextInput focuses the next input field
func (a *authModel) nextInput() {
	a.focused = (a.focused + 1) % len(a.inputs)
}

// prevInput focuses the previous input field
func (a *authModel) prevInput() {
	a.focused--
	// Wrap around
	if a.focused < 0 {
		a.focused = len(a.inputs) - 1
	}
}

func Login(cfg *config.Config) error {
	clientID, err := keyring.Get(cfg.Service, "CLIENT_ID")
	if err != nil {
		log.Error("error getting client id from keyring", "error", err)
	}

	clientSecret, err := keyring.Get(cfg.Service, "CLIENT_SECRET")
	if err != nil {
		log.Error("error getting client secret from keyring", "error", err)
	}

	// get any empty client keys through bubbletea
	if clientID != "" && clientSecret != "" {
		cfg.ClientID = clientID
		cfg.ClientSecret = clientSecret

		return nil
	}

	p := tea.NewProgram(newAuthModel())
	if _, err := p.Run(); err != nil {
		log.Error("failed to launch tui", "error", err)
	}

	return nil
}
