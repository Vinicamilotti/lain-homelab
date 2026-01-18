package frontend

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"gitub.com/Vinicamilotti/homelab-lain-utils-servicecli/models"
)

type ConfigServicesMenu struct {
	Config    models.ConfigFile
	Selection []string                     // items on the to-do list
	Cursor    int                          // which to-do list item our cursor is pointing at
	Selected  map[int]models.ServiceConfig // which to-do items are selected
	Save      bool
}

func configModel(config models.ConfigFile) ConfigServicesMenu {
	services := make([]string, len(config.AutoStart))
	choices := make(map[int]models.ServiceConfig, len(config.AutoStart))
	i := 0
	for k, v := range config.AutoStart {
		services[i] = k

		if v.Tracked {
			choices[i] = v
		}
		i++
	}
	return ConfigServicesMenu{
		Config: config,
		// Our to-do list is a grocery list
		Selection: services,
		Selected:  choices,
	}
}

func (m ConfigServicesMenu) Init() tea.Cmd {
	return nil
}
func (m ConfigServicesMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.Cursor < len(m.Selection)-1 {
				m.Cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = m.Config.AutoStart[m.Selection[m.Cursor]]
			}

		case "s":
			m.Save = true
			return m, tea.Quit
		}

	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m ConfigServicesMenu) View() string {
	// The header
	s := "Choose what programs to track\n\n"

	// Iterate over our choices
	for i, choice := range m.Selection {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.Cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.Selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"
	s += "Press s to save configuration.\n"

	// Send the UI for rendering
	return s
}
