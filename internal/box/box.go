package box

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Border(lipgloss.NormalBorder(), true, true, true, true).
		BorderForeground(lipgloss.Color("#b8bb26"))
)

type Bubble struct {
	height int
	width  int
}

func New(width int, height int) Bubble {
	return Bubble{
		width:  width,
		height: height,
	}
}

func (b *Bubble) SetWidth(width int) {
	b.width = width - style.GetHorizontalFrameSize()
}

func (b Bubble) Init() tea.Cmd {
	return nil
}

func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	return b, nil
}

func (b Bubble) View() string {
	return style.Width(b.width).Height(b.height).Render("")
}
