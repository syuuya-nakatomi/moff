package utils

import (
	"strings"

	"github.com/rivo/tview"
)

func ToggleCheckbox(checkbox *tview.Checkbox) {
	checkbox.SetChecked(!checkbox.IsChecked())
}

func GetSelectedItems(checkboxList *tview.CheckboxList) []string {
	var selectedItems []string

	for _, item := range checkboxList.GetItems() {
		if checkboxList.GetChecked(item) {
			selectedItems = append(selectedItems, strings.Split(item, " ")[0])
		}
	}

	return selectedItems
}

func FormatSelectedItems(selectedItems []string) string {
	return strings.Join(selectedItems, "\n- ")
}
