package main

import (
	"fmt"
	"github.com/Sabooboo/pokecli/cmd"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	// TODO Remove explicit Setenv
	os.Setenv("DEBUG", "yes")
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	// Default command is ui, so running app runs all the fancy tui stuff :-)
	cmd.Execute()
}
