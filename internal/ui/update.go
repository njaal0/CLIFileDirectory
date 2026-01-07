package ui

import (
	"fmt"
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

		maxOffset := len(m.entries) - m.viewportHeight
		if maxOffset < 0 {
			maxOffset = 0
		}

		if m.scrollOffset > maxOffset {
			m.scrollOffset = maxOffset
		}

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			fmt.Println(m.currentPath)
			return m, tea.Quit

		case "up":
			if m.selectedIdx > 0 {
				m.selectedIdx--
				if m.selectedIdx < m.scrollOffset {
					m.scrollOffset = m.selectedIdx
				}
			}

		case "down":
			if m.selectedIdx < len(m.entries)-1 {
				m.selectedIdx++
				if m.selectedIdx >= m.scrollOffset+m.viewportHeight {
					m.scrollOffset++
				}
			}

		case "right":
			if len(m.entries) == 0 {
				return m, nil
			}

			entry := m.entries[m.selectedIdx]

			if !entry.IsDir() {
				return m, nil
			}

			fullPath := filepath.Join(m.currentPath, entry.Name())
			newEntries, err := fs.ListEntries(fullPath)

			if err != nil {
				return m, nil
			}

			m.history = append(m.history, m.currentPath)
			m.currentPath = fullPath
			m.entries = newEntries
			m.selectedIdx = 0
			m.scrollOffset = 0

		case "left":
			var targetPath string

			if len(m.history) > 0 {
				targetPath = m.history[len(m.history)-1]
				m.history = m.history[:len(m.history)-1]
			} else {
				parent := filepath.Dir(m.currentPath)

				if parent == m.currentPath {
					return m, nil
				}

				targetPath = parent

			}

			newEntries, err := fs.ListEntries(targetPath)
			if err != nil {
				return m, nil
			}

			m.currentPath = targetPath
			m.entries = newEntries
			m.selectedIdx = 0
			m.scrollOffset = 0
		}

	}

	return m, nil
}
