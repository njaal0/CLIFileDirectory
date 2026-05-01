package ui

import (
	"fmt"
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

	start := m.ScrollOffset

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

	if m.Searching {
		matchInfo := footerStyle.Render("  " + fmt.Sprintf("(%d matches)", len(m.Entries)))
		prompt := inputLabelStyle.Render("  Search: ") + inputStyle.Render(m.SearchQuery+"█") + matchInfo
		b.WriteString(prompt + "\n")
	} else if m.CreatingFolder {
		prompt := inputLabelStyle.Render("  New folder: ") + inputStyle.Render(m.NewFolderName+"█")
		b.WriteString(prompt + "\n")
	} else if m.Renaming {
		prompt := inputLabelStyle.Render("  Rename to: ") + inputStyle.Render(m.RenameTo+"█")
		b.WriteString(prompt + "\n")
	} else {
		b.WriteString(footerStyle.Render("  ↑↓ navigate   → open   ← back   / search   n new folder   r rename   q quit"))
	}

	return b.String()
}
