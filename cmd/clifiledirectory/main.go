package clifiledirectory

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	startPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory", err)
		os.Exit(1)
	}

	model := ui.newModel(startPath)

	program := tea.NewProgram(model)
	if err := program.Start(); err != nil {
		fmt.Println("Error starting program", err)
		os.Exit(1)
	}

}
