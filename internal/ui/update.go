package ui

import (
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/njaal0/CLIFileDirectory/internal/fs"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.viewportHeight = msg.Height - 6
		if m.viewportHeight < 1 {
			m.viewportHeight = 1
		}
		m.viewportWidth = msg.Width

		maxOffset := len(m.Entries) - m.viewportHeight
		if maxOffset < 0 {
			maxOffset = 0
		}

		if m.ScrollOffset > maxOffset {
			m.ScrollOffset = maxOffset
		}

	case tea.KeyMsg:
		// While in search mode, intercept runes/backspace; pass navigation through.
		if m.Searching {
			switch msg.Type {
			case tea.KeyEsc:
				m.Searching = false
				m.SearchQuery = ""
				m.Entries = m.AllEntries
				m.SelectedIdx = 0
				m.ScrollOffset = 0

			case tea.KeyBackspace, tea.KeyDelete:
				if len(m.SearchQuery) > 0 {
					m.SearchQuery = m.SearchQuery[:len(m.SearchQuery)-1]
					m.Entries = filterEntries(m.AllEntries, m.SearchQuery)
					m.SelectedIdx = 0
					m.ScrollOffset = 0
				}

			case tea.KeyRunes:
				m.SearchQuery += string(msg.Runes)
				m.Entries = filterEntries(m.AllEntries, m.SearchQuery)
				m.SelectedIdx = 0
				m.ScrollOffset = 0

			case tea.KeyUp:
				if m.SelectedIdx > 0 {
					m.SelectedIdx--
					if m.SelectedIdx < m.ScrollOffset {
						m.ScrollOffset = m.SelectedIdx
					}
				}

			case tea.KeyDown:
				if m.SelectedIdx < len(m.Entries)-1 {
					m.SelectedIdx++
					if m.SelectedIdx >= m.ScrollOffset+m.viewportHeight {
						m.ScrollOffset++
					}
				}

			case tea.KeyRight:
				if len(m.Entries) == 0 {
					break
				}
				entry := m.Entries[m.SelectedIdx]
				if !entry.IsDir() {
					break
				}
				fullPath := filepath.Join(m.CurrentPath, entry.Name())
				newEntries, err := fs.ListEntries(fullPath)
				if err != nil {
					break
				}
				m.History = append(m.History, m.CurrentPath)
				m.CurrentPath = fullPath
				m.AllEntries = newEntries
				m.Entries = newEntries
				m.Searching = false
				m.SearchQuery = ""
				m.SelectedIdx = 0
				m.ScrollOffset = 0
			}
			return m, nil
		}

		// While in rename mode, intercept all keys for the input.
		if m.Renaming {
			switch msg.Type {
			case tea.KeyEsc:
				m.Renaming = false
				m.RenameTo = ""

			case tea.KeyEnter:
				if m.RenameTo != "" && len(m.Entries) > 0 {
					oldPath := filepath.Join(m.CurrentPath, m.Entries[m.SelectedIdx].Name())
					newPath := filepath.Join(m.CurrentPath, m.RenameTo)
					if err := fs.RenameEntry(oldPath, newPath); err == nil {
						if entries, err := fs.ListEntries(m.CurrentPath); err == nil {
							m.Entries = entries
						}
					}
				}
				m.Renaming = false
				m.RenameTo = ""

			case tea.KeyBackspace, tea.KeyDelete:
				if len(m.RenameTo) > 0 {
					m.RenameTo = m.RenameTo[:len(m.RenameTo)-1]
				}

			case tea.KeyRunes:
				m.RenameTo += string(msg.Runes)
			}
			return m, nil
		}

		// While in folder-creation mode, intercept all keys for the input.
		if m.CreatingFolder {
			switch msg.Type {
			case tea.KeyEsc:
				m.CreatingFolder = false
				m.NewFolderName = ""

			case tea.KeyEnter:
				if m.NewFolderName != "" {
					newPath := filepath.Join(m.CurrentPath, m.NewFolderName)
					if err := fs.CreateDir(newPath); err == nil {
						if entries, err := fs.ListEntries(m.CurrentPath); err == nil {
							m.Entries = entries
						}
					}
				}
				m.CreatingFolder = false
				m.NewFolderName = ""

			case tea.KeyBackspace, tea.KeyDelete:
				if len(m.NewFolderName) > 0 {
					m.NewFolderName = m.NewFolderName[:len(m.NewFolderName)-1]
				}

			case tea.KeyRunes:
				m.NewFolderName += string(msg.Runes)
			}
			return m, nil
		}

		switch msg.String() {

		case "c":
			m.ShouldPrintPath = false

		case "q":
			m.ShouldPrintPath = true
			return m, tea.Quit

		case "n":
			m.CreatingFolder = true
			m.NewFolderName = ""

		case "/":
			m.Searching = true
			m.SearchQuery = ""
			m.Entries = m.AllEntries
			m.SelectedIdx = 0
			m.ScrollOffset = 0

		case "r":
			if len(m.Entries) > 0 {
				m.Renaming = true
				m.RenameTo = m.Entries[m.SelectedIdx].Name()
			}

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
			m.AllEntries = newEntries
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
			m.AllEntries = newEntries
			m.Entries = newEntries
			m.SelectedIdx = 0
			m.ScrollOffset = 0
		}

	}

	return m, nil
}

func filterEntries(entries []os.DirEntry, query string) []os.DirEntry {
	if query == "" {
		return entries
	}
	q := strings.ToLower(query)
	var result []os.DirEntry
	for _, e := range entries {
		if strings.Contains(strings.ToLower(e.Name()), q) {
			result = append(result, e)
		}
	}
	if result == nil {
		return []os.DirEntry{}
	}
	return result
}
