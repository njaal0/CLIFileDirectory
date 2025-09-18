package clifiledirectory

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	startPath = os.Getwd()
	model = ui.newModel(startPath)

	program = tea.NewProgram(model)
	error = program.Run()
	if error != nil {
		print(error)
		os.Exit(1)
	}
}
