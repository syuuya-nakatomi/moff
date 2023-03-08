package ui

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func SelectVulnerabilities(vulnerabilities []string) ([]string, error) {
	// Initialize the TUI
	if err := termui.Init(); err != nil {
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
	events := termui.PollEvents()
	for {
		select {
		case e := <-events:
			switch e.ID {
			case "q", "<C-c>":
				// Quit if user presses 'q' or Ctrl+C
				return nil, nil
			case "<Enter>":
				// Add selected vulnerability to list
				i := list.SelectedRow
				if i >= 0 {
					selected = append(selected, vulnerabilities[i])
				}
			case "<Up>", "<Down>":
				// Move selection up or down
				dir := 1
				if e.ID == "<Up>" {
					dir = -1
				}
				i := list.SelectedRow
				if i >= 0 {
					i += dir
					if i >= len(vulnerabilities) {
						i = len(vulnerabilities) - 1
					} else if i < 0 {
						i = 0
					}
					list.SelectedRow = i
				}
			}
		}

		// Update the list widget
		list.SelectedRowStyle = termui.NewStyle(termui.ColorClear, termui.ColorGreen)
		list.SelectedRow = len(selected)
		termui.Render(grid)
	}
}
