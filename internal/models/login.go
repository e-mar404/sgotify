package models

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	cid = iota
	cs
)

type loginModel struct {
	inputs []textinput.Model
	focused int
}

func NewLoginModel() loginModel {
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
		inputs:  inputs,
		focused: 0,
	}
}

func (a loginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (a loginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(a.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if a.focused == len(a.inputs) - 1 {
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

	for i := range a.inputs {
		a.inputs[i], cmds[i] = a.inputs[i].Update(msg)
	}
	return a, tea.Batch(cmds...)
}

func (a loginModel) View() string {
	buf := strings.Builder{}

	buf.WriteString("Enter Spotify Client ID:\n")
	buf.WriteString(a.inputs[cid].View() + "\n\n")
	buf.WriteString("Enter Spotify Client Secret:\n")
	buf.WriteString(a.inputs[cs].View() + "\n")
	buf.WriteString("continue ->")

	return buf.String()
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
