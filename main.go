package main

import (
	"github.com/charmbracelet/lipgloss"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"tea-demo/internal/box"
)

type Bubble struct {
	left  box.Bubble
	right box.Bubble
}

func (b Bubble) Init() tea.Cmd {
	return nil
}

func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		width := msg.Width
		b.left.SetWidth(width / 2)
		b.right.SetWidth(width / 2)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return b, tea.Quit
		}
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd

	b.left, cmd = b.left.Update(msg)
	cmds = append(cmds, cmd)

	b.right, cmd = b.right.Update(msg)
	cmds = append(cmds, cmd)

	return b, nil
}

func (b Bubble) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		b.left.View(),
		b.right.View(),
	)
}

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			panic("unable to create debug file")
		}
		defer f.Close()
	}

	b := Bubble{
		left:  box.New(20, 15),
		right: box.New(20, 15),
	}

	if _, err := tea.NewProgram(b, tea.WithAltScreen()).Run(); err != nil {
		panic(err)
	}
}
