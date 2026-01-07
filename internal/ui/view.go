package ui

import (
	"fmt"
)

func (m *Model) View() string {
	output := fmt.Sprintf("📁 %s\n\n", m.currentPath)

	if len(m.entries) == 0 {
		return output + "(empty directory)\n"
	}

	start := m.scrollOffset
	end := start + m.viewportHeight
	if end > len(m.entries) {
		end = len(m.entries)
	}

	for i := start; i < end; i++ {
		entry := m.entries[i]

		cursor := "  "
		if i == m.selectedIdx {
			cursor = "➜ "
		}

		name := entry.Name()
		if entry.IsDir() {
			name += "/"
		}

		output += fmt.Sprintf("%s%s\n", cursor, name)
	}

	return output
}
