package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/njaal0/CLIFileDirectory/internal/fs"
)

type Model struct {
	CurrentPath     string
	Entries         []os.DirEntry
	SelectedIdx     int
	History         []string
	viewportHeight  int
	viewportWidth   int
	ScrollOffset    int
	ShouldPrintPath bool
	CreatingFolder  bool
	NewFolderName   string
	Renaming        bool
	RenameTo        string
}

func NewModel(startPath string) *Model {
	entries, err := fs.ListEntries(startPath)
	if err != nil {
		fmt.Println("Error listing entries", err)
		entries = []os.DirEntry{}
	}

	return &Model{
		CurrentPath:     startPath,
		Entries:         entries,
		SelectedIdx:     0,
		History:         []string{},
		viewportHeight:  15,
		viewportWidth:   80,
		ShouldPrintPath: false,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}
