package ui

import (
	"strings"
)

func (m *Model) View() string {
	var b strings.Builder

	// Header
	b.WriteString(headerStyle.Render("📁 " + m.CurrentPath))
	b.WriteString("\n\n")

	if len(m.Entries) == 0 {
		b.WriteString("(empty directory)\n")
		return b.String()
	}

	start := m.ScrollOffset
	end := start + m.viewportHeight
	if end > len(m.Entries) {
		end = len(m.Entries)
	}

	for i := start; i < end; i++ {
		entry := m.Entries[i]

		name := entry.Name()
		style := fileStyle

		if entry.IsDir() {
			name += "/"
			style = dirStyle
		}

		var line string
		if i == m.SelectedIdx {
			line = selectedStyle.Render("➜ " + name)
		} else {
			line = style.Render("  " + name)
		}

		b.WriteString(line + "\n")
	}

	// Footer
	b.WriteString("\n")
	b.WriteString(footerStyle.Render("↑ ↓: navigate   ➜: open   ←: backtrack   q: quit   c: cancel cd after quit"))

	return b.String()
}
