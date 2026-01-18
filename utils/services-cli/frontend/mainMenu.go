package frontend

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type MainMenu struct {
	cursor int // which item our cursor is on
	Choice int // which item was selected
}

func mainMenuModel() MainMenu {
	return MainMenu{cursor: 1}
}

func (m MainMenu) Init() tea.Cmd {
	return nil
}

func (m MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < 4 {
				m.cursor++
			}
		case "enter":
			m.Choice = m.cursor
			return m, tea.Quit
		case "1", "2", "3", "4", "5":
			m.Choice = int(msg.String()[0] - '1')
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MainMenu) View() string {
	s := "Choose an option (1-5):\n\n"
	options := make(map[int]string, 0)

	options[0] = "Configure Services"
	options[1] = "List Tracked Services"
	options[2] = "Restart Tracked Services"
	options[3] = "Stop Tracked Services"
	options[4] = "Exit"

	for i := 0; i <= 4; i++ {
		cursor := " "
		iStr := fmt.Sprint(i + 1)
		if m.cursor == i {
			cursor = ">" // Highlight the current selection
		}
		opt := options[i]
		s += fmt.Sprintf("%s %s - %s\n", cursor, iStr, opt)
	}

	s += "\n(Press up/down to move, enter to select, q to quit)\n"
	return s
}
