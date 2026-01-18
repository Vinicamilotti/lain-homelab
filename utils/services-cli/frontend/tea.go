package frontend

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"gitub.com/Vinicamilotti/homelab-lain-utils-servicecli/models"
)

func StartMainMenu() MainMenu {
	p := tea.NewProgram(mainMenuModel())
	var m tea.Model
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	if resp, ok := m.(MainMenu); ok {
		return resp
	}

	panic("could not type assert model")

}

func StartConfigUI(config models.ConfigFile) ConfigServicesMenu {
	p := tea.NewProgram(configModel(config))
	var m tea.Model
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	if resp, ok := m.(ConfigServicesMenu); ok {
		return resp
	}

	panic("could not type assert model")
}
