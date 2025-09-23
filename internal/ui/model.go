package ui

import (
	"fmt"
	"github.com/njaal0/CLIFileDirectory/internal/fs"
	"os"
)

type Model struct {
	currentPath string
	entries     []os.DirEntry
	selectedIdx int
	history     []string
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
