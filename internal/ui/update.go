package ui

import (
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/njaal0/CLIFileDirectory/internal/fs"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "q":
			return m, tea.Quit

		case "c":
			m.ShouldPrintPath = false

		case "up":
			if m.SelectedIdx > 0 {
				m.SelectedIdx--
			}
			if m.SelectedIdx < m.ScrollOffset {
				m.ScrollOffset = m.SelectedIdx
			}

		case "down":
			if m.SelectedIdx < len(m.Entries)-1 {
				m.SelectedIdx++
			}
			if m.SelectedIdx >= m.ScrollOffset+m.viewportHeight {
				m.ScrollOffset++
			}

		case "enter":
			entry := m.Entries[m.SelectedIdx]
			if !entry.IsDir() {
				return m, nil
			}

			m.History = append(m.History, m.CurrentPath)

			nextPath := filepath.Join(m.CurrentPath, entry.Name())
			entries, err := fs.ListEntries(nextPath)
			if err != nil {
				return m, nil
			}

			m.CurrentPath = nextPath
			m.Entries = entries
			m.SelectedIdx = 0
			m.ScrollOffset = 0
		}
	}

	return m, nil
}
