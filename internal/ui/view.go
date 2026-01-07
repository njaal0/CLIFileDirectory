package ui

func (m *Model) View() string {
	output := headerStyle.Render("📁 " + m.currentPath)
	output += "\n\n"

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

		name := entry.Name()
		style := fileStyle

		if entry.IsDir() {
			name += "/"
			style = dirStyle
		}

		line := "  " + name

		if i == m.selectedIdx {
			line = selectedStyle.Render("➜ " + name)
		} else {
			line = style.Render(line)
		}

		output += line + "\n"
	}

	output += "\n"
	output += footerStyle.Render("↑ ↓: navigate   ➜: open   ←: back   q: quit")

	return output
}
