package constants

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	P          *tea.Program
	WindowSize tea.WindowSizeMsg
)

const (
	Purple = lipgloss.Color("189")
	Gray   = lipgloss.Color("245")
	Green  = lipgloss.Color("42")
)

var WindowStyle = lipgloss.NewStyle().
	Align(lipgloss.Center, lipgloss.Center).
	Border(lipgloss.RoundedBorder())

var SecondaryTextCLI = lipgloss.NewStyle().
	Foreground(Green)

var (
	HeaderStyle = lipgloss.NewStyle().Foreground(Purple).Bold(true).Align(lipgloss.Center)
	CellStyle   = lipgloss.NewStyle().Padding(0, 1).Width(14)
	RowStyle    = CellStyle.Foreground(Gray)
	BorderStyle = lipgloss.NewStyle().Foreground(Purple)
)

var HelpTableCLI = table.New().
	Border(lipgloss.ThickBorder()).
	BorderStyle(BorderStyle).
	StyleFunc(func(row, col int) lipgloss.Style {
		var style lipgloss.Style

		switch row {
		case table.HeaderRow:
			return HeaderStyle
		default:
			style = RowStyle
		}

		// Make the second column a little wider.
		if col == 1 {
			style = style.Width(22)
		}

		return style
	})
