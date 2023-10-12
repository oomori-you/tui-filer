/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

type Flex struct {
	Header *tview.TextView
	Body   *tview.List
	Footer *tview.InputField
}

// cdCmd represents the cd command
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()

		wdPath, _ := os.Getwd()
		textview := tview.NewTextView()
		textview.SetText(wdPath).SetTitle("Top").SetBorder(true)

		inputField := tview.NewInputField()
		inputField.SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})

		list := tview.NewList()
		list.ShowSecondaryText(false).SetBorder(true).SetTitle("Middle Proportion")

		f := new(Flex)
		f.Header = textview
		f.Body = list
		f.Footer = inputField

		f.ConstructList()
		f.Body.SetInputCapture(f.CaptureList)

		flex := tview.NewFlex().
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(f.Header, 3, 0, false).
				AddItem(f.Body, 0, 1, true).
				AddItem(f.Footer, 1, 0, false),
				0, 1, true)

		if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	},
}

func (f Flex) CaptureList(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEnter:
		index := f.Body.GetCurrentItem()
		main, _ := f.Body.GetItemText(index)
		f.Header.SetText(f.Header.GetText(true) + "/" + main)
		f.ConstructList()
		return nil
	}

	switch event.Rune() {
	case 'h':
		return tcell.NewEventKey(tcell.KeyLeft, ' ', tcell.ModNone)
	case 'j':
		return tcell.NewEventKey(tcell.KeyDown, ' ', tcell.ModNone)
	case 'k':
		return tcell.NewEventKey(tcell.KeyUp, ' ', tcell.ModNone)
	case 'l':
		return tcell.NewEventKey(tcell.KeyRight, ' ', tcell.ModNone)
	}

	return event // 上記以外のキー入力をdefaultのキーアクションへ伝える
}

func (f Flex) ConstructList() {
	f.Body.Clear()
	files, _ := os.ReadDir(f.Header.GetText(true))
	for _, v := range files {
		f.Body.AddItem(v.Name(), "", 0, nil)
	}
}

func init() {
	rootCmd.AddCommand(cdCmd)
}
