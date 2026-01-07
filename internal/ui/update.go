package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/njaal0/CLIFileDirectory/internal/fs"
	"path/filepath"
	//"os"
	//"path/filepath"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.viewportHeight = msg.Height - 3
		if m.viewportHeight < 1 {
			m.viewportHeight = 1
		}

		maxOffset := len(m.Entries) - m.viewportHeight
		if maxOffset < 0 {
			maxOffset = 0
		}

		if m.ScrollOffset > maxOffset {
			m.ScrollOffset = maxOffset
		}

	case tea.KeyMsg:
		switch msg.String() {

		case "c":
			m.ShouldPrintPath = false

		case "q":
			return m, tea.Quit

		case "up":
			if m.SelectedIdx > 0 {
				m.SelectedIdx--
				if m.SelectedIdx < m.ScrollOffset {
					m.ScrollOffset = m.SelectedIdx
				}
			}

		case "down":
			if m.SelectedIdx < len(m.Entries)-1 {
				m.SelectedIdx++
				if m.SelectedIdx >= m.ScrollOffset+m.viewportHeight {
					m.ScrollOffset++
				}
			}

		case "right":
			if len(m.Entries) == 0 {
				return m, nil
			}

			entry := m.Entries[m.SelectedIdx]

			if !entry.IsDir() {
				return m, nil
			}

			fullPath := filepath.Join(m.CurrentPath, entry.Name())
			newEntries, err := fs.ListEntries(fullPath)

			if err != nil {
				return m, nil
			}

			m.History = append(m.History, m.CurrentPath)
			m.CurrentPath = fullPath
			m.Entries = newEntries
			m.SelectedIdx = 0
			m.ScrollOffset = 0

		case "left":
			var targetPath string

			if len(m.History) > 0 {
				targetPath = m.History[len(m.History)-1]
				m.History = m.History[:len(m.History)-1]
			} else {
				parent := filepath.Dir(m.CurrentPath)

				if parent == m.CurrentPath {
					return m, nil
				}

				targetPath = parent

			}

			newEntries, err := fs.ListEntries(targetPath)
			if err != nil {
				return m, nil
			}

			m.CurrentPath = targetPath
			m.Entries = newEntries
			m.SelectedIdx = 0
			m.ScrollOffset = 0
		}

	}

	return m, nil
}
