package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// リストを作成
	buttons := tview.NewList()
	buttons.AddItem("Yes", "", rune(0), nil)
	buttons.AddItem("No", "", rune(0), nil)

	// テキストを表示するためのViewを作成
	selectionText := tview.NewBox().SetBorder(true).SetTitle("Selection")

	// ボタンを選択したときの動作を設定
	buttons.SetSelectedFunc(func(index int, label string, primaryText string, secondaryText rune) {
		selectionText.SetTitle(label)
	})

	// レイアウトを作成
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(buttons, 0, 3, false).
		AddItem(selectionText, 0, 1, false)

	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}
