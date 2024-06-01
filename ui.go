package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	Ready   = 0
	Playing = 1
)

type Status int

type model struct {
	station Station
	status  Status
}

func initialModel(s Station) model {
	return model{
		station: s,
		status:  Ready,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	content := ""

	if m.status == Ready {
		content = "Press [Space] to stream"
	} else if m.status == Playing {
		content = fmt.Sprintf("Now Playing: %s", m.station.title)
	}

	return content
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			fmt.Println("ctrl")
			return m, tea.Quit
		case " ":
			if m.status == Ready {
				done := make(chan bool)
				go Play(m.station.streamURL, done)
				m.status = Playing
			} else if m.status == Playing {
				Stop()
				m.status = Ready
			}

			return m, nil
		}
	}

	return m, nil
}
