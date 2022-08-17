package common

import tea "github.com/charmbracelet/bubbletea"

type Component interface {
	tea.Model
	SetSize(width, height int) Component
}
