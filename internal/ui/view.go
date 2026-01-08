package ui

import (
	"strings"
)

func (m *Model) View() string {
	var b strings.Builder

	b.WriteString("📁 " + m.CurrentPath + "\n\n")

	start := 0
	if m.SelectedIdx >= m.viewportHeight {
		start = m.SelectedIdx - m.viewportHeight + 1
	}

	end := start + m.viewportHeight
	if end > len(m.Entries) {
		end = len(m.Entries)
	}

	for i := start; i < end; i++ {
		entry := m.Entries[i]
		name := entry.Name()

		if entry.IsDir() {
			name += "/"
		}

		if i == m.SelectedIdx {
			b.WriteString("> " + name + "\n")
		} else {
			b.WriteString("  " + name + "\n")
		}
	}

	for i := end - start; i < m.viewportHeight; i++ {
		b.WriteString("\n")
	}

	b.WriteString("\n↑↓: navigate   →: open   ←: back   c: cancel cd after quit   q: quit")

	return b.String()
}
