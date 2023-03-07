package tui

import (
	"fmt"

	"github.com/myuser/myapp/internal"
	"github.com/myuser/myapp/pkg/utils"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
)

type TUI struct {
	app          *tview.Application
	checkboxList *tview.CheckboxList
}

func NewTUI() *TUI {
	app := tview.NewApplication()

	checkboxList := tview.NewCheckboxList().
		SetTitle("Select items to deploy").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetBorderColor(tview.Styles.BorderColor)

	button := tview.NewButton("Generate").
		SetSelectedFunc(func() {
			selectedItems := utils.GetSelectedItems(checkboxList)

			if len(selectedItems) == 0 {
				logrus.Warn("No items selected")
				return
			}

			playbook, err := internal.GeneratePlaybook(selectedItems)
			if err != nil {
				logrus.WithError(err).Error("Failed to generate playbook")
				return
			}

			fmt.Println(playbook)
		})

	// Add the button to the flex layout
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(checkboxList, 0, 1, true).
		AddItem(button, 1, 0, false)

	return &TUI{
		app:          app,
		checkboxList: checkboxList,
	}
}

func (t *TUI) Run() error {
	vulsResults, err := internal.ParseVulsScanResults("vuls-results.json")
	if err != nil {
		return err
	}

	for _, result := range vulsResults {
		item := result.Hostname
		for _, vuln := range result.Vulnerabilities {
			item += fmt.Sprintf(" (%s)", vuln.ID)
		}
		t.checkboxList.AddItem(item, "", 0, nil)
	}

	if err := t.app.SetRoot(t.checkboxList, true).SetFocus(t.checkboxList).Run(); err != nil {
		return err
	}

	return nil
}
