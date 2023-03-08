package ui

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func SelectVulnerabilities(vulnerabilities []string) ([]string, error) {
	// Initialize the TUI
	err := termui.Init()
	if err != nil {
		return nil, err
	}
	defer termui.Close()

	// Create the list widget and populate it with vulnerabilities
	list := widgets.NewList()
	list.Title = "Select vulnerabilities to deploy"
	list.Rows = vulnerabilities

	// Create the layout and add the list widget
	grid := termui.NewGrid()
	grid.Set(
		termui.NewRow(1.0,
			termui.NewCol(1.0, list),
		),
	)
	termui.Render(grid)

	// Wait for user input
	var selected []string
	for {
		e := termui.PollEvent()
		if e.Type == termui.KeyboardEvent {
			switch e.ID {
			case "q", "<C-c>":
				// Quit if user presses 'q' or Ctrl+C
				return nil, nil
			case "<Enter>":
				// Add selected vulnerability to list
				_, i := list.Selected()
				if i >= 0 {
					selected = append(selected, vulnerabilities[i])
				}
			}
		}

		// Update the list widget
		list.SelectedRowStyle = termui.NewStyle(termui.ColorClear, termui.ColorGreen)
		list.SelectedRow = len(selected)
		termui.Render(grid)
	}
}
