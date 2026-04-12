package ui

import (
	"os"

	"github.com/charmbracelet/lipgloss"
)

// renderer is bound to stderr — the actual terminal output — so that lipgloss
// correctly detects color support even when stdout is piped (e.g. via a shell
// wrapper that captures the selected path).
var renderer = lipgloss.NewRenderer(os.Stderr)

var (
	headerStyle = renderer.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("63"))

	selectedStyle = renderer.NewStyle().
			Foreground(lipgloss.Color("230")).
			Background(lipgloss.Color("62")).
			Bold(true)

	dirStyle = renderer.NewStyle().
			Foreground(lipgloss.Color("69")).
			Bold(true)

	fileStyle = renderer.NewStyle().
			Foreground(lipgloss.Color("252"))

	footerStyle = renderer.NewStyle().
			Foreground(lipgloss.Color("241"))

	separatorStyle = renderer.NewStyle().
			Foreground(lipgloss.Color("237"))
)
