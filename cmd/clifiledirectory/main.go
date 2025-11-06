package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/njaal0/CLIFileDirectory/internal/ui"
	"os"
)

func main() {
	startPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory", err)
		os.Exit(1)
	}

	model := ui.NewModel(startPath)
	program := tea.NewProgram(model)

	_, err = program.Run()
	if err != nil {
		fmt.Println("Error starting program", err)
		os.Exit(1)
	}
}
