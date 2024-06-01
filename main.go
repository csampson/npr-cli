package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	station := Station{
		title:     "CPR",
		streamURL: "https://stream1.cprnetwork.org/cpr1_lo",
	}

	m := initialModel(station)
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
