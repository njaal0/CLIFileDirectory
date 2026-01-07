package ui

import (
	"strings"
)

func (m *Model) View() string {
	var b strings.Builder

	b.WriteString("📁 " + m.CurrentPath + "\n\n")

	if len(m.Entries) == 0 {
		b.WriteString("(empty)\n")
		return b.String()
	}

	start := m.ScrollOffset
	end := start + m.viewportHeight
	if end > len(m.Entries) {
		end = len(m.Entries)
	}

	for i := start; i < end; i++ {
		prefix := "  "
		if i == m.SelectedIdx {
			prefix = "➜ "
		}
		b.WriteString(prefix + m.Entries[i].Name() + "\n")
	}

	b.WriteString("↑ ↓: navigate   ➜: open   ←: back   q: quit   c: cancel cd after quit")
	return b.String()
}
