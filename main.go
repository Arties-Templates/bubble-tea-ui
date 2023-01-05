package main

import "fmt"
import "os"
import tea "github.com/charmbracelet/bubbletea"

type model struct{ state string }

func initialModel() model { return model{state: "Nothing here!!"} }
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrlc", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	vw := "Hello, World!\n\n"

	vw += "Press q To Quit."

	return vw
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Err: %s", err)
		os.Exit(1)
	}
}