package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/njaal0/CLIFileDirectory/internal/fs"
	"os"
	"path/filepath"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.selectedIdx > 0 {
				m.selectedIdx--
			}

		case "down":
			if m.selectedIdx < len(m.entries)-1 {
				m.selectedIdx++
			}

		case "enter":
			if len(m.entries) == 0 {
				return m, nil
			}

			entry := m.entries[m.selectedIdx]
			if entry.IsDir() {
				fullPath := filepath.Join(m.currentPath, entry.Name())
				newEntries, err := fs.ListEntries(fullPath)
				if err != nil {
					return m, nil
				}
				m.history = append(m.entries, os.DirEntry(m.currentPath))
				m.currentPath = fullPath
				m.entries = newEntries
				m.selectedIdx = 0
			}

		case "backspace":
			if len(m.history) > 0 {
				return m, nil
			}

			previousPath := m.history[len(m.history)-1]
			newEntries, err := fs.ListEntries(previousPath)
			if err != nil {
				return m, nil
			}
			m.history = m.history[:len(m.history)-1]
			m.entries = newEntries
			m.selectedIdx = 0
		}

	}

	return m, nil
}
