package ui

import (
	"strings"
)

func (m *Model) View() string {
	if m.ShouldPrintPath {
		return ""
	}

	width := m.viewportWidth
	if width == 0 {
		width = 80
	}

	sep := separatorStyle.Render(strings.Repeat("─", width))

	var b strings.Builder

	b.WriteString(headerStyle.Render("  📁 "+m.CurrentPath) + "\n")
	b.WriteString(sep + "\n\n")

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

		var line string
		if i == m.SelectedIdx {
			line = selectedStyle.Width(width).Render("▶ " + name)
		} else if entry.IsDir() {
			line = dirStyle.Render("  " + name)
		} else {
			line = fileStyle.Render("  " + name)
		}
		b.WriteString(line + "\n")
	}

	for i := end - start; i < m.viewportHeight; i++ {
		b.WriteString("\n")
	}

	b.WriteString("\n" + sep + "\n")
	b.WriteString(footerStyle.Render("  ↑↓ navigate   → open   ← back   q quit"))

	return b.String()
}
