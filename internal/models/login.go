package models

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zalando/go-keyring"
)

const service = "sgotify-auth"
const (
	cid = iota
	cs
)

type sessionState int
const (
	promptView = iota
	inputView
	authView
)

type loginModel struct {
	state sessionState
	inputs []textinput.Model
	focused int
}

func NewLoginModel() loginModel {
	// Need to check if theyre are pre existing credentials. If there are saved ask if user wasnt to use those or not. If there are no creds send to input screen
	var state sessionState
	_, foundCID := keyring.Get(service, "CLIENT_ID")
	_, foundCS := keyring.Get(service, "CLIENT_SECRET")
	if foundCID == nil && foundCS == nil {
		state = promptView 
	} else {
		state = inputView 
	}

	var inputs []textinput.Model = make([]textinput.Model, 2)
	inputs[cid] = textinput.New()
	inputs[cid].Placeholder = "CLIENT_ID"
	inputs[cid].Focus()
	inputs[cid].Width = 50
	inputs[cid].Prompt = ""

	inputs[cs] = textinput.New()
	inputs[cs].Placeholder = "CLIENT_SECRET"
	inputs[cs].Width = 50
	inputs[cs].Prompt = ""

	return loginModel{
		state: state,
		inputs:  inputs,
		focused: 0,
	}
}

func (a loginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (a loginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(a.inputs))

	switch a.state {
	case promptView:
		switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String(){
					case "y":
						a.state = authView
						return a, nil

					case "n":
						a.state = inputView
						return a, nil

					case "ctrl+c", "esc":
						return a, tea.Quit
				}

			case error:
				return a, nil
		}

	case inputView:
		switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.Type {
					case tea.KeyEnter:
						if a.focused == len(a.inputs) - 1 {
							keyring.Set(service, "CLIENT_ID", a.inputs[cid].Value())
							keyring.Set(service, "CLIENT_SECRET", a.inputs[cs].Value())

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

			case error:
				return a, nil
		}
	case authView:
		switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.Type {
					case tea.KeyCtrlC, tea.KeyEsc:
						return a, tea.Quit
			}
		}
	}

	for i := range a.inputs {
		a.inputs[i], cmds[i] = a.inputs[i].Update(msg)
	}
	return a, tea.Batch(cmds...)
}

func (a loginModel) View() string {
	switch a.state {
	case promptView:
		return "There are already some credentials saved do you want to use those? [Y|N]"
	case inputView:
		buf := strings.Builder{}

		buf.WriteString("Enter Spotify Client ID:\n")
		buf.WriteString(a.inputs[cid].View() + "\n\n")
		buf.WriteString("Enter Spotify Client Secret:\n")
		buf.WriteString(a.inputs[cs].View() + "\n")
		buf.WriteString("continue ->")

		return buf.String()
	case authView:
		return "auth view"
	}

	return "something went terribly wrong"
} 

func (a *loginModel) nextInput() {
	a.focused = (a.focused + 1) % len(a.inputs)
	a.focusInput(a.focused)
}

func (a *loginModel) prevInput() {
	a.focused--
	if a.focused < 0 {
		a.focused = len(a.inputs) - 1
	}

	a.focusInput(a.focused)
}

func (a *loginModel) focusInput(i int) {
	for i := range a.inputs {
		a.inputs[i].Blur()
	}
	a.inputs[i].Focus()
}
