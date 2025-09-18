package ui

import "fmt"

type Model struct {
	currentPath string
	entries     []DirEntry
	selectedIdx int
	history     []string
}

func NewModel(startPath string) *Model {
	entries, err := fs.ListEntries(startPath)
	if err != nil {
		fmt.Println("Error listing entries", err)
		entries = []DirEntry{}
	}

	return model{
		currentPath: startPath,
		entries:     entries,
		selectedIdx: 0,
		history:     []string{},
	}
}
