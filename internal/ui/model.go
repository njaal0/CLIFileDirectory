package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/njaal0/CLIFileDirectory/internal/fs"
	"os"
)

type Model struct {
	currentPath    string
	entries        []os.DirEntry
	selectedIdx    int
	history        []string
	viewportHeight int
	scrollOffset   int
}

func NewModel(startPath string) *Model {
	entries, err := fs.ListEntries(startPath)
	if err != nil {
		fmt.Println("Error listing entries", err)
		entries = []os.DirEntry{}
	}

	return &Model{
		currentPath: startPath,
		entries:     entries,
		selectedIdx: 0,
		history:     []string{},
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}
