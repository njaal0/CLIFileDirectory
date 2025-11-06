package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/njaal0/CLIFileDirectory/internal/fs"
	"os"
)

type Model struct {
	currentPath string
	entries     []os.DirEntry
	selectedIdx int
	history     []string
}

func NewModel(startPath string) *Model {
	entries, err := fs.ListEntries(startPath)
	if err != nil {
		fmt.Println("Error listing entries", err)
		entries = []os.DirEntry{}
	}

	return &Model{
		currentPath: startPath,
		entries:     entries,
		selectedIdx: 0,
		history:     []string{},
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) View() string {
	if len(m.entries) == 0 {
		return fmt.Sprintf("Current directory: %s\n\n(empty)\n", m.currentPath)
	}

	s := fmt.Sprintf("Current directory: %s\n\n", m.currentPath)

	for i, entry := range m.entries {
		prefix := "   "
		if i == m.selectedIdx {
			prefix = "-> "
		}

		name := entry.Name()

		if entry.IsDir() {
			name += "/"
		}

		s += fmt.Sprintf("%s%s\n", prefix, name)
	}

	return s
}
