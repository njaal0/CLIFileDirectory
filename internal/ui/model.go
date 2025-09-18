package ui

import (
	"fmt"
	"github.com/njaal0/CLIFileDirectory/internal/fs"
)

type Model struct {
	currentPath string
	entries     []DirEntry
	selectedIdx int
	history     []string
}

func newModel(startPath string) *Model {
	entries, err := fs.listEntries(startPath)
	if err != nil {
		fmt.Println("Error listing entries", err)
		entries = []DirEntry{}
	}

	return &Model{
		currentPath: startPath,
		entries:     entries,
		selectedIdx: 0,
		history:     []string{},
	}
}
